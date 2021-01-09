package redis_sortdset

import (
	"framework/class/sortedset"
	"github.com/go-redis/redis"
)

type Options struct {
	redisClient *redis.Client
}

type Option func(o *Options)

func New(opts ...Option) sortedset.Sortedset {
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

