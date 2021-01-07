package stream_redis_stream_v8

import (
	"framework/class/logger"
	"framework/class/stream"
	"github.com/go-redis/redis/v8"
)

type Options struct {
	redisClient *redis.Client
	logger      logger.Logger
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

func WithLogger(log logger.Logger) Option {
	return func(o *Options) {
		o.logger = log
	}
}
