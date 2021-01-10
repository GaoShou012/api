package im_v1

import (
	"encoding/json"
	"fmt"
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

// 加密消息
// messageType 0x01 消息，0x02 通知
func (p *plugin) encodeMessage(messageType uint8, message []byte) ([]byte, error) {
	max := len(message)
	res := make([]byte, len(message)+1)
	res[0] = messageType
	for i := 0; i < max; i++ {
		res[i+1] = message[i]
	}
	return res, nil
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

func (p *plugin) CreateChannel(topic string, info channel.Info) error {
	return p.opts.channel.Create(topic, info)
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

	if message == nil {
		err = fmt.Errorf("消息不能为空")
		return
	}

	encodedMessage := make([]byte, len(message)+1)
	encodedMessage[0] = 0x01
	for i := 0; i < len(message); i++ {
		encodedMessage[i+1] = message[i]
	}
	encodedMessage, err = p.encodeMessage(eventTypeMessage, message)
	if err != nil {
		return
	}

	messageId, err = p.opts.client.Push(uuid, encodedMessage)
	if err != nil {
		return
	}
	err = p.opts.gateway.Publish(uuid, "notice", nil)
	if err != nil {
		return
	}
	return
}

func (p *plugin) clientEvent(messageId string, data []byte) (*clientEvent, error) {
	evt := &clientEvent{
		id:   messageId,
		data: nil,
	}
	switch data[0] {
	case eventTypeMessage:
		evt.data = data[1:]
		break
	case eventTypeChannelNotice:
		// 从频道消息流拉取消息
		notice := &clientEventOfChannelNotice{}
		if err := json.Unmarshal(data[1:], notice); err != nil {
			return nil, err
		}
		msg, err := p.opts.channel.PullById(notice.topic, notice.messageId)
		if err != nil {
			return nil, err
		}

		// 等于消息ID，使用客户端的消息流
		// 消息数据，使用channel的消息
		evt.data = msg
		break
	}
	return evt, nil
}

func (p *plugin) PullMessageFromClient(uuid string, lastMessageId string, count uint64) ([]client.Event, error) {
	res, err := p.opts.client.Pull(uuid, lastMessageId, count)
	if err != nil {
		return nil, err
	}

	events := make([]client.Event, 0)

	for _, row := range res {
		evt, err := p.clientEvent(row.Id(), row.Data())
		if err != nil {
			env.Logger.Error(err, uuid, row.Id(), row.Data())
			continue
		}
		events = append(events, evt)
	}

	return events, nil
}

func (p *plugin) PullMessageFromClientById(uuid string, messageId string) (client.Event, error) {
	event, err := p.opts.client.PullById(uuid, messageId)
	if err != nil {
		return nil, err
	}

	return p.clientEvent(event.Id(), event.Data())
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
