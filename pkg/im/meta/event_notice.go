package meta

import "im/class/client"

type EventDataOfNotice struct {
	Topic string
	MsgId string
}

func (e *EventDataOfNotice) Type() client.EventType {
	return client.EventTypeNotice
}
