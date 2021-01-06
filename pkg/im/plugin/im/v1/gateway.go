package im_v1

import (
	"encoding/json"
	"fmt"
	"im/class/channel"
	"im/class/client"
	"im/class/im"
	"im/env"
	"im/meta"
)

type gatewayEventType int

const (
	gatewayEventTypeClientAttach gatewayEventType = iota
	gatewayEventTypeClientDetach
	gatewayEventTypeClientNotice
	gatewayEventTypeClientMessage
)

type gatewayEventClientAttach struct {
	UUID          string
	LastMessageId string
}
type gatewayEventClientDetach struct {
	UUID string
}
type gatewayEventClientNotice struct {
	UUID string
}

type gatewayEvent struct {
	Type gatewayEventType
	Data interface{}
}

type gatewayHandler struct {
	channel.Channel
	client.Client
	im.OnPublishCallback

	clients  map[string]string
	onEvents chan *gatewayEvent
}

func (g *gatewayHandler) Init() {
	g.clients = make(map[string]string)
	g.onEvents = make(chan *gatewayEvent, 10000)
}

func (g *gatewayHandler) ClientAttach(uuid string, lastMessageId string) {
	g.onEvents <- &gatewayEvent{
		Type: gatewayEventTypeClientAttach,
		Data: &gatewayEventClientAttach{
			UUID:          uuid,
			LastMessageId: lastMessageId,
		},
	}
}

func (g *gatewayHandler) ClientDetach(uuid string) {
	g.onEvents <- &gatewayEvent{
		Type: gatewayEventTypeClientDetach,
		Data: &gatewayEventClientDetach{UUID: uuid},
	}
}

func (g *gatewayHandler) ClientNotice(uuid string) {
	g.onEvents <- &gatewayEvent{
		Type: gatewayEventTypeClientNotice,
		Data: &gatewayEventClientNotice{UUID: uuid},
	}
}

func (g *gatewayHandler) Run() {
	clients := make(map[string]string)
	gatewayEvents := make(chan *gatewayEvent, 100000)
	go func() {
		for {
			gatewayEvt := <-gatewayEvents

			switch gatewayEvt.Type {
			case gatewayEventTypeClientAttach:
				evt := gatewayEvt.Data.(*gatewayEventClientAttach)
				clients[evt.UUID] = evt.LastMessageId
				// 投放到历史记录同步器

				break
			case gatewayEventTypeClientDetach:
				evt := gatewayEvt.Data.(*gatewayEventClientDetach)
				delete(clients, evt.UUID)
				break
			case gatewayEventTypeClientNotice:
				evt := gatewayEvt.Data.(*gatewayEventClientNotice)
				clientUUID := evt.UUID
				lastMessageId, ok := clients[evt.UUID]
				if !ok {
					continue
				}

				events, err := g.Client.Pull(clientUUID, lastMessageId, 5)
				if err != nil {
					env.Logger.Error(err)
					continue
				}
				for _, event := range events {
					// update the last message id of the client
					clients[clientUUID] = event.Id()

					// decode client event
					clientEvt := &meta.ClientEvent{}
					if err := json.Unmarshal(event.Data(), clientEvt); err != nil {
						env.Logger.Error(err)
						continue
					}

					/*
						channel type:
						Get the message form the channel then proxy to client.
						message type:
						Direct message proxy to client.
					*/
					switch clientEvt.Type {
					case meta.ClientEventTypeChannel:
						topic := clientEvt.ClientEventChannel.Topic
						msgId := clientEvt.ClientEventChannel.MsgId
						data, err := g.Channel.PullById(topic, msgId)
						if err != nil {
							env.Logger.Error(err)
							break
						}
						g.OnPublishCallback(clientUUID, data)
						break
					case meta.ClientEventTypeMessage:
						g.OnPublishCallback(clientUUID, clientEvt.ClientEventMessage.Payload)
						break
					default:
						desc := fmt.Errorf("Unknown client event\n")
						env.Logger.Error(desc)
						break
					}
				}
				break
			default:
				desc := fmt.Errorf("Unknown gateway event,gatewayEvt=%v\n", gatewayEvt)
				env.Logger.Error(desc)
				break
			}
		}
	}()
}
