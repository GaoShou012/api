package stream_redis_stream

import (
	"framework/class/stream"
	"github.com/go-redis/redis/v7"
)

var _ stream.Event = &Event{}

type Event struct {
	id          string
	streamName  string
	message     []byte
	redisClient *redis.ClusterClient
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
	_, err := e.redisClient.XDel(e.streamName, e.id).Result()
	return err
}

func (e *Event) Error() error {
	return e.err
}
