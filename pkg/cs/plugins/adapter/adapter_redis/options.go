package adapter_redis

import (
	"cs/class"
	"github.com/go-redis/redis/v8"
)

type Option func(o *Options)

type Options struct {
	redisClient *redis.Client
}

func NewAdapter(redisClient *redis.Client, opts ...Option) class.Adapter {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	n := &Adapter{
		redisClient: redisClient,
		opts:        nil,
	}

	if err := n.Init(); err != nil {
		panic(err)
	}

	return n
}

func WithRedisClient(redisClient *redis.Client) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}
