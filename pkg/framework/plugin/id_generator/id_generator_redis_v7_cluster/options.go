package id_generator_redis_v7_cluster

import (
	"framework/class/id_generator"
	"github.com/go-redis/redis/v7"
)

type Options struct {
	redisClient *redis.ClusterClient
}

type Option func(o *Options)

func New(opts ...Option) id_generator.IdGenerator {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	p := &plugin{
		opts: options,
	}
	if err := p.Init(); err != nil {
		panic(err)
	}

	return p
}

func WithRedisClient(redisClient *redis.ClusterClient) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}
