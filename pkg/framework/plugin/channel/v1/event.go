package channel_v1

import (
	"encoding/json"
	"framework/class/channel"
)

type Event struct {
	T channel.EventType
	D interface{}
}

func (e *Event) Encode() ([]byte,error) {
	return json.Marshal(e)
}