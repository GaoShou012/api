package v1

import (
	"api/cs/class/client_event"
	"api/cs/class/session"
	"cs/env"
	"encoding/json"
	"fmt"
)

func topicOfSessionRecords(sessionId string) string {
	return fmt.Sprintf("stream:session:records:%s", sessionId)
}
func topicOfClientEvents(uuid string) string {
	return fmt.Sprintf("stream:client:events:%s", uuid)
}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) GetAllClients(sessionId string) ([]*session.ClientItem, error) {
	return nil, nil
}

func (p *plugin) Broadcast(sessionId string, message []byte) error {
	var notice *client_event.Notice

	// 消息压入会话stream
	{
		topic := topicOfSessionRecords(sessionId)
		msgId, err := p.opts.Stream.Push(topic, message)
		if err != nil {
			return err
		}
		notice = &client_event.Notice{
			Topic: topic,
			MsgId: msgId,
		}
	}

	// 获取所有成员列表，进行通知
	{
		clients, err := p.GetAllClients(sessionId)
		if err != nil {
			return err
		}

		if len(clients) == 0 {
			return nil
		}

		event := &client_event.Event{}
		event.Type = client_event.EventTypeNotice
		event.Notice = notice
		j, err := json.Marshal(event)
		if err != nil {
			return err
		}

		for _, client := range clients {
			// 产生一个消息通知
			topic := topicOfClientEvents(client.UUID)

			if _, err := p.opts.Stream.Push(topic, j); err != nil {
				desc := fmt.Sprintf("通知失败:topic=%s,clientUUID=%s,err=%v", notice.Topic, client.UUID, err)
				env.Logger.Error(desc)
			}

			// 推送客户端事件
			if err := env.Gateway.PublishClientEvent(client.UUID, j); err != nil {
				desc := fmt.Sprintf("推送失败:err=%v", err)
				env.Logger.Error(desc)
			}
		}
	}

	return nil
}
