package im_v1

import (
	"fmt"
	"im/class/client"
	"im/class/im"
	"im/env"
	"im/meta"
	"sync"
	"time"
)

const (
	NumberOfRecordsPerPulling = 20
	CGTime                    = time.Hour
	PullTime                  = time.Millisecond * 250
)

type SyncRecord struct {
	client.Client
	im.OnPublishCallback

	mutex   sync.Mutex
	clients map[string]string
	events  chan *syncRecordEvent
}

func (s *SyncRecord) AddClient(uuid string, lastMessageId string) {
	s.events <- &syncRecordEvent{
		Type: syncRecordEventTypeClientAttach,
		Data: &syncRecordEventClientAttach{
			UUID:          uuid,
			LastMessageId: lastMessageId,
		},
	}
}
func (s *SyncRecord) DelClient(uuid string) {
	s.events <- &syncRecordEvent{
		Type: syncRecordEventTypeClientDetach,
		Data: &syncRecordEventClientDetach{UUID: uuid},
	}
}

func (s *SyncRecord) Init() {
	s.events = make(chan *syncRecordEvent, 10000)
}

func (s *SyncRecord) Run() {
	go func() {
		cg := time.NewTicker(CGTime)
		pull := time.NewTicker(PullTime)
		for {
			select {
			case event := <-s.events:
				switch event.Type {
				case syncRecordEventTypeClientAttach:
					evt := event.Data.(*syncRecordEventClientAttach)
					if _, ok := s.clients[evt.UUID]; ok {
						desc := fmt.Errorf("同步记录器，客户端已经存在uuid=%s", evt.UUID)
						env.Logger.Warn(desc)
						continue
					}
					s.clients[evt.UUID] = evt.LastMessageId
					break
				case syncRecordEventTypeClientDetach:
					evt := event.Data.(*syncRecordEventClientDetach)
					if _, ok := s.clients[evt.UUID]; !ok {
						desc := fmt.Errorf("同步记录器，客户端不存在，不能detach,uuid=%s", evt.UUID)
						env.Logger.Warn(desc)
						continue
					}
					delete(s.clients, evt.UUID)
					break
				}
			case <-pull.C:
				pull.Stop()
				s.Pull()
				pull.Reset(PullTime)
			case <-cg.C:
				tmp := make(map[string]string)
				for key, val := range s.clients {
					tmp[key] = val
				}
				s.clients = tmp
			}
		}
	}()
}

func (s *SyncRecord) Pull() {
	for key, val := range s.clients {
		clientUUID, clientMessageId := key, val
		events, err := s.Client.Pull(clientUUID, clientMessageId, NumberOfRecordsPerPulling)
		if err != nil {
			env.Logger.Error(err)
			continue
		}

		for _, event := range events {
			newId := event.Id()
			if meta.IsNewMessageId(clientMessageId, newId) == false {
				continue
			}
			clientMessageId = newId
			s.OnPublishCallback(clientUUID, event.Data())
		}

		if len(events) < NumberOfRecordsPerPulling {
			delete(s.clients, clientUUID)
		}
	}
}
