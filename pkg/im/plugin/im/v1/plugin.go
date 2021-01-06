package im_v1

import (
	"encoding/json"
	"im/class/channel"
	"im/class/client"
	"im/class/gateway"
	"im/class/im"
	"im/env"
	"im/meta"
)

var _ im.IM = &plugin{}

type plugin struct {
	syncRecord      *SyncRecord
	gatewayHandlers *gatewayHandler
	opts            *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Channel() channel.Channel {
	return p.opts.channel
}

func (p *plugin) Client() client.Client {
	return p.opts.client
}

func (p *plugin) Gateway() gateway.Gateway {
	return p.opts.gateway
}

func (p *plugin) ClientAttach(uuid string, lastMessageId string) {
	p.syncRecord.AddClient(uuid, lastMessageId)
}

func (p *plugin) ClientDetach(uuid string) {
	p.syncRecord.DelClient(uuid)
}

func (p *plugin) CreateChannel(info channel.Info) error {
	return p.opts.channel.Create(info)
}

func (p *plugin) DeleteChannel(topic string) error {
	return p.opts.channel.Delete(topic)
}

func (p *plugin) ClientsOfChannel(topic string) (channel.Clients, error) {
	return p.opts.channel.Clients(topic)
}

func (p *plugin) ChannelsOfClient(clientUUID string) (client.Channels, error) {
	return p.opts.client.Channels(clientUUID)
}

func (p *plugin) PushMessageToChannel(topic string, message []byte) (messageId string, err error) {
	var clients channel.Clients
	clients, err = p.opts.channel.Clients(topic)
	if err != nil {
		return
	}

	return p.PushMessageToChannelWithClients(topic, message, clients)
}

func (p *plugin) PushMessageToChannelWithClients(topic string, message []byte, clients channel.Clients) (messageId string, err error) {
	// 推送消息到频道
	messageId, err = p.opts.channel.Push(topic, message)
	if err != nil {
		return
	}

	// 推送通知到客户端事件流
	var event []byte
	{
		evt := &meta.ClientEvent{
			Type: meta.ClientEventTypeChannel,
			ClientEventChannel: &meta.ClientEventChannel{
				Topic: topic,
				MsgId: messageId,
			},
			ClientEventMessage: nil,
		}
		event, err = json.Marshal(evt)
		if err != nil {
			return
		}
		count := 0
		arr := make([]string, len(clients))
		for uuid, _ := range clients {
			arr[count] = uuid
			count++
		}
		if err := p.opts.client.PushClients(arr, event); err != nil {
			env.Logger.Error(err)
		}
	}

	// 推送通知到网关
	for uuid, _ := range clients {
		if err := p.opts.gateway.Publish(uuid, "notice", nil); err != nil {
			env.Logger.Error(err)
			continue
		}
	}

	return
}

func (p *plugin) PushMessageToClient(uuid string, message []byte) error {
	// 推送消息到网关
	return p.opts.gateway.Publish(uuid, "message", message)
}

func (p *plugin) PushMessageToClientEvent(uuid string, message []byte) (messageId string, err error) {
	// 推送消息到客户端的消息流
	// 推送客户端通知到网关
	messageId, err = p.opts.client.Push(uuid, message)
	if err != nil {
		return
	}
	err = p.opts.gateway.Publish(uuid, "notice", nil)
	if err != nil {
		return
	}
	return
}

func (p *plugin) ClientSubscribeChannel(uuid string, topic string) error {
	if err := p.opts.client.Subscribe(uuid, topic); err != nil {
		return err
	}
	if err := p.opts.channel.Subscribe(topic, uuid); err != nil {
		return err
	}
	return nil
}

func (p *plugin) ClientUnSubscribeChannel(uuid string, topic string) error {
	if err := p.opts.client.UnSubscribe(uuid, topic); err != nil {
		return err
	}
	if err := p.opts.channel.UnSubscribe(topic, uuid); err != nil {
		return err
	}
	return nil
}

func (p *plugin) OnPublish(callback im.OnPublishCallback) {
	{
		handler := &gatewayHandler{
			Channel:           p.opts.channel,
			Client:            p.opts.client,
			OnPublishCallback: callback,
			clients:           nil,
			onEvents:          nil,
		}
		handler.Init()
		handler.Run()
		p.gatewayHandlers = handler
	}
	p.gatewayHandlers = &gatewayHandler{}

	p.opts.gateway.Subscribe(func(message gateway.Message) error {
		header := message.Header()
		uuid := header["UUID"]
		messageType := header["MessageType"]
		switch messageType {
		case "notice":
			// 客户端有新的消息通知
			p.gatewayHandlers.ClientNotice(uuid)
			break
		case "message":
			// 直接推送消息
			callback(uuid, message.Body())
			break
		}
		return nil
	})
}
