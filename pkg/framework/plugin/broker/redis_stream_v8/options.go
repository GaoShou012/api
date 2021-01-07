package broker_redis_stream_v8

import (
	"framework/class/broker"
	"github.com/go-redis/redis/v8"
)

type Options struct {
	redisClient *redis.Client
}

type Option func(o *Options)

func New(opts ...Option) broker.Broker {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	b := &plugin{
		redisClient: nil,
		opts:        options,
	}
	if err := b.Init(); err != nil {
		panic(err)
	}

	return b
}

func WithRedisClient(redisClient *redis.Client) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}
