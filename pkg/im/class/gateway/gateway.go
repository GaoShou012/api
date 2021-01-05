package gateway

import "framework/class/broker"

type MessageType int8

const (
	MessageTypeControl = iota
	MessageTypeText
)

type Gateway interface {
	Init() error
	Publish(uuid string, message []byte) error
	PublishControl(control Control) error
	Subscribe(handler Handler) (broker.Subscriber, error)
}

type Handler func(message Message) error

type Message interface {
	Type() MessageType
	Header() map[string]string
	Body() []byte
}

type Control interface {
	CtlType() string
}
