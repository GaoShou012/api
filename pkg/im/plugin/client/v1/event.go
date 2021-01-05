package client_v1

import (
	"encoding/json"
	"im/class/client"
)

type event struct {
	T client.EventType
	D interface{}
}

func (e *event) Type() client.EventType {
	return e.T
}
func (e *event) Data() interface{} {
	return e.D
}

func EncodeEvent(eventData client.EventData) ([]byte, error) {
	evt := &event{
		T: eventData.Type(),
		D: eventData,
	}
	return json.Marshal(evt)
}
