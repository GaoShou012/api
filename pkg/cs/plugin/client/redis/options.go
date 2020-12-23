package client_redis

import (
	"cs/class/client"
	"github.com/go-redis/redis/v8"
)

type Options struct {
	redisClient *redis.Client
}

type Option func(o *Options)

func New(opts ...Option) client.Client {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	p := &plugin{
		redisClient: nil,
		opts:        options,
	}
	if err := p.Init(); err != nil {
		panic(err)
	}

	return p
}

func WithRedisClient(redisClient *redis.Client) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}

//func WithLogger(log logger.Logger) Option {
//	return func(o *Options) {
//		o.logger = log
//	}
//}
