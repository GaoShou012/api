package stream_redis_stream

import (
	"framework/class/stream"
	"github.com/go-redis/redis"
)

type Options struct {
	redisClient *redis.Client
}

type Option func(o *Options)

func New(opts ...Option) stream.Stream {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	s := &plugin{
		opts:        options,
	}
	if err := s.Init(); err != nil {
		panic(err)
	}

	return s
}

func WithRedisClient(redisClient *redis.Client) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}

