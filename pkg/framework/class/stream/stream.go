package stream

type Stream interface {
	Init() error
	Connect(dns string) error
	Delete(topic string) error
	Push(topic string, message []byte) (string, error)
	Pull(topic string, lastMessageId string, count uint64) ([]Event, error)
	RevPull(topic string, lastMessageId string, count uint64) ([]Event, error)
	PullById(topic string, messageId string) (Event, error)
	Subscribe(topic string, handler Handler) (Subscriber, error)
}

type Handler func(evt Event) error

type Event interface {
	Id() string
	Topic() string
	Message() []byte
	Ack() error
	Error() error
}
type Subscriber interface {
	UnSubscribe() error
}
