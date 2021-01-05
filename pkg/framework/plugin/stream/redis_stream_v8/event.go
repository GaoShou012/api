package stream_redis_stream_v8

import (
	"context"
	"framework/class/stream"
	"github.com/go-redis/redis/v8"
)

var _ stream.Event = &Event{}

type Event struct {
	id          string
	streamName  string
	message     []byte
	redisClient *redis.Client
	err         error
}

func (e *Event) Id() string {
	return e.id
}

func (e *Event) Topic() string {
	return e.streamName
}

func (e *Event) Message() []byte {
	return e.message
}

func (e *Event) Ack() error {
	_, err := e.redisClient.XDel(context.TODO(), e.streamName, e.id).Result()
	return err
}

func (e *Event) Error() error {
	return e.err
}
