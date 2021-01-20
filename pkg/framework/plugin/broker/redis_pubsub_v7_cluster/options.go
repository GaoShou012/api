package broker_redis_pubsub_v7_cluster

import (
	"framework/class/broker"
	"framework/class/logger"
	"framework/env"
	"github.com/go-redis/redis/v7"
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

	if options.logger == nil {
		options.logger = env.Logger
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

func WithLogger(log logger.Logger) Option {
	return func(o *Options) {
		o.logger = log
	}
}
