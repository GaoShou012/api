package broker

type Handler func(evt Event) error

type Broker interface {
	Init() error
	Connect(dns string) error
	Publish(topic string, message []byte) error
	Subscribe(topic string, handler Handler) (Subscriber, error)
}

type Event interface {
	Topic() string
	Header() map[string]string
	Message() []byte
	Ack() error
	Error() error
}

type Subscriber interface {
	UnSubscribe() error
}
