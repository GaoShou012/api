package stream_redis_stream

import (
	"context"
	"framework/class/stream"
	"github.com/go-redis/redis/v8"
)

var _ stream.Event = &redisEvent{}

type redisEvent struct {
	id          string
	streamName  string
	message     []byte
	redisClient *redis.Client
	err         error
}

func (e *redisEvent) Id() string {
	return e.id
}

func (e *redisEvent) Topic() string {
	return e.streamName
}

func (e *redisEvent) Message() []byte {
	return e.message
}

func (e *redisEvent) Ack() error {
	_, err := e.redisClient.XDel(context.TODO(), e.streamName, e.id).Result()
	return err
}

func (e *redisEvent) Error() error {
	return e.err
}
