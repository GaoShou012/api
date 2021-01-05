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
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) GetChannelAdapter() channel.Channel {
	return p.opts.channelAdapter
}

func (p *plugin) GetClientAdapter() client.Client {
	return p.opts.clientAdapter
}

func (p *plugin) CreateChannel(info channel.Info) error {
	return p.opts.channelAdapter.Create(info)
}

func (p *plugin) DeleteChannel(topic string) error {
	return p.opts.channelAdapter.Delete(topic)
}

func (p *plugin) ClientsOfChannel(topic string) (channel.Clients, error) {
	return p.opts.channelAdapter.Clients(topic)
}

func (p *plugin) ChannelsOfClient(clientUUID string) (client.Channels, error) {
	return p.opts.clientAdapter.Channels(clientUUID)
}

func (p *plugin) Publish(topic string, message []byte) (messageId string, err error) {
	var clients channel.Clients
	clients, err = p.opts.channelAdapter.Clients(topic)
	if err != nil {
		return
	}

	return p.PublishWithClients(topic, message, clients)
}

func (p *plugin) PublishWithClients(topic string, message []byte, clients channel.Clients) (messageId string, err error) {
	// 推送消息到频道
	messageId, err = p.opts.channelAdapter.Publish(topic, message)
	if err != nil {
		return
	}

	// 推送通知到客户端事件流
	{
		eventData := &meta.EventDataOfNotice{
			Topic: topic,
			MsgId: messageId,
		}
		count := 0
		arr := make([]string, len(clients))
		for uuid, _ := range clients {
			arr[count] = uuid
			count++
		}
		if err := p.opts.clientAdapter.PushClients(arr, eventData); err != nil {
			env.Logger.Error(err)
		}
	}

	// 推送通知到网关
	for uuid, _ := range clients {
		if err := p.opts.gateway.Publish(uuid, []byte("event")); err != nil {
			env.Logger.Error(err)
			continue
		}
	}

	return
}

func (p *plugin) Subscribe(topic string, clientUUID string) error {
	if err := p.opts.channelAdapter.Subscribe(topic, clientUUID); err != nil {
		return err
	}

	if err := p.opts.clientAdapter.Subscribe(clientUUID, topic); err != nil {
		return err
	}
	return nil
}

func (p *plugin) UnSubscribe(topic string, clientUUID string) error {
	if err := p.opts.channelAdapter.UnSubscribe(topic, clientUUID); err != nil {
		return err
	}

	if err := p.opts.clientAdapter.UnSubscribe(clientUUID, topic); err != nil {
		return err
	}
	return nil
}

func (p *plugin) RunGateway(handler im.GatewayHandler) {
	p.opts.gateway.Subscribe(func(message gateway.Message) error {

		evt := &meta.Event{}
		if err := evt.Decode(message.Body()); err != nil {
			desc := fmt.Errorf("解析事件失败,data=%v", message.Body())
			env.Logger.Error(desc)
			return nil
		}

		clientUUID := message.Header()["UUID"]

		if evt.Type == meta.EventTypeNotice {
			events, err := p.opts.client.Pull(clientUUID, 3)
			if err != nil {
				env.Logger.Error(err)
				return nil
			}
			for _, event := range events {
				evt := &clientEvent{}
				if err := json.Unmarshal(event, evt); err != nil {
					env.Logger.Error(err)
					continue
				}
				switch evt.Type {
				case clientEventTypeNotice:
					p.opts.client.Pull(clientUUID, 3)
					break
				case clientEventTypeMessage:
					handler(evt.Data)
					break
				}
			}
		} else {
			handler(evt.Data)
		}

		return nil
	})
}
