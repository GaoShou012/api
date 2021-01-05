package cs

import (
	"encoding/json"
	"fmt"
	"framework/class/stream"
)

type EventType uint64

const (
	EventTypeNotification = iota
)

type Notification struct {
	Topic string
	MsgId string
}

type Event struct {
	Type EventType
	*Notification
}

type PersonalEvent struct {
	stream.Stream
}

func (p *PersonalEvent) Topic(uuid string) string {
	return fmt.Sprintf("cs:personal_event_stream:%s", uuid)
}

func (p *PersonalEvent) Push(uuid string, event *Event) (string, error) {
	topic := p.Topic(uuid)
	j, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	return p.Stream.Push(topic, j)
}
func (p *PersonalEvent) Pull(uuid string, lastMessageId string, count uint64) ([]*Event, error) {
	topic := p.Topic(uuid)
	res, err := p.Stream.Pull(topic, lastMessageId, count)
	if err != nil {
		return nil, err
	}
	events := make([]*Event, 0)
	for _, row := range res {
		msg := row.Message()
		event := &Event{}
		if err := json.Unmarshal(msg, event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
