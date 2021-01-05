package meta

import "encoding/json"

type EventType int

const (
	// 如果事件是Notice
	// Data = 客户端uuid
	// 如果事件是Message
	// Data = 消息内容
	EventTypeNotice EventType = iota
	EventTypeMessage
)

type Event struct {
	Type EventType
	Data []byte
}

func (e *Event) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Event) Decode(data []byte) error {
	return json.Unmarshal(data, e)
}
