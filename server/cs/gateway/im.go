package gateway

import (
	"api/cs/env"
	"im/class/client"
	"im/class/gateway"
	"im/meta"
)

func Handler() {
	env.IM.GetGateway().Subscribe(func(message gateway.Message) error {
		switch message.Type() {
		case gateway.MessageTypeText:
			// decode message
			evt := &meta.Event{}
			if err := evt.Decode(message.Body()); err != nil {
				return err
			}

			switch evt.Type {
			case meta.EventTypeNotice:
				// TODO 拉取消息通知
				cli := env.IM.GetClient()
				events, err := cli.Pull(string(evt.Data), 10)
				if err != nil {
					return err
				}
				for _, e := range events {
					switch e.Type() {
					case client.EventTypeNotice:
						break
					case client.EventTypeMessage:
						break
					}
				}
				break
			case meta.EventTypeMessage:
				// TODO 推送消息给ws
				break
			}
			break
		case gateway.MessageTypeControl:
			break
		}
		return nil
	})
}
