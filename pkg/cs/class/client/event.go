package client

type EventType int

const (
	// 消息，直接转发给客户端
	EventTypeMessage EventType = iota
	// 通知，需要读取源topic数据进行转发
	EventTypeNotice
)

type Event interface {
	Type() EventType
	Data() interface{}
}
