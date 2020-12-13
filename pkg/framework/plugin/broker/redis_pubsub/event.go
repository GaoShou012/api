package broker_redis_pubsub

import (
	"framework/class/broker"
)

/*
	Event
*/
var _ broker.Event = &redisPubSubEvent{}

type redisPubSubEvent struct {
	topic   string
	message []byte
	err     error
}

func (e *redisPubSubEvent) Header() map[string]string {
	return nil
}

func (e *redisPubSubEvent) Topic() string {
	return e.topic
}

func (e *redisPubSubEvent) Message() []byte {
	return e.message
}

func (e *redisPubSubEvent) Ack() error {
	return nil
}

func (e *redisPubSubEvent) Error() error {
	return e.err
}
