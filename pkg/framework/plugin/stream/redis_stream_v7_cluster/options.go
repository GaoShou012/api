package stream_redis_stream

import (
	"framework/class/stream"
	"github.com/go-redis/redis/v7"
)

type Options struct {
	redisClient *redis.ClusterClient
}

type Option func(o *Options)

func NewStream(opts ...Option) stream.Stream {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	s := &plugin{
		opts: options,
	}
	if err := s.Init(); err != nil {
		panic(err)
	}

	return s
}

func WithRedisClient(redisClient *redis.ClusterClient) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}
