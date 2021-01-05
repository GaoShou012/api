package client_event

type EventType int

const (
	// 消息，直接转发给客户端
	EventTypeMessage EventType = iota
	// 通知，需要读取源topic数据进行转发
	EventTypeNotice
)

type Notice struct {
	Topic string
	MsgId string
}
type Message struct {
	Payload []byte
}

type Event struct {
	Type EventType
	*Notice
	*Message
}

type ClientEvent interface {
	Push(uuid string, event *Event) error
	Pull(uuid string, lastMessageId string, count uint64) ([]*Event, error)
}
