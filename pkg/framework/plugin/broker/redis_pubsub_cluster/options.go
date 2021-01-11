package broker_redis_pubsub_cluster

import (
	"framework/class/broker"
	"framework/class/logger"
	"github.com/go-redis/redis"
)

type Options struct {
	redisClient *redis.ClusterClient
	logger      logger.Logger
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

func WithRedisClient(redisClient *redis.ClusterClient) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}
