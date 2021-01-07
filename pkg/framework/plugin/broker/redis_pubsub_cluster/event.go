package broker_redis_pubsub_cluster

import (
	"framework/class/broker"
)

/*
	Event
*/
var _ broker.Event = &event{}

type event struct {
	topic   string
	message []byte
	err     error
}

func (e *event) Header() map[string]string {
	return nil
}

func (e *event) Topic() string {
	return e.topic
}

func (e *event) Message() []byte {
	return e.message
}

func (e *event) Ack() error {
	return nil
}

func (e *event) Error() error {
	return e.err
}
