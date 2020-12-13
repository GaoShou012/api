package broker_redis_stream

import (
	"context"
	"framework/class/broker"
	"github.com/go-redis/redis/v8"
)

var _ broker.Event = &redisStreamEvent{}

type redisStreamEvent struct {
	header      map[string]string
	topic       string
	message     []byte
	redisClient *redis.Client
	err         error
}

func (e *redisStreamEvent) Header() map[string]string {
	return e.header
}

func (e *redisStreamEvent) Topic() string {
	return e.topic
}

func (e *redisStreamEvent) Message() []byte {
	return e.message
}

func (e *redisStreamEvent) Ack() error {
	id := e.header["Id"]
	_, err := e.redisClient.XDel(context.TODO(), e.topic, id).Result()
	return err
}

func (e *redisStreamEvent) Error() error {
	return e.err
}
