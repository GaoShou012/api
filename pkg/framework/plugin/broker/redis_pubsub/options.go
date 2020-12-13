package broker_redis_pubsub

import (
	"framework/class/broker"
	"framework/class/logger"
	"github.com/go-redis/redis/v8"
)

type Options struct {
	redisClient *redis.Client
	logger      logger.Logger
}

type Option func(o *Options)

func NewBroker(opts ...Option) broker.Broker {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	b := &redisPubSub{
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

func WithLogger(log logger.Logger) Option {
	return func(o *Options) {
		o.logger = log
	}
}
