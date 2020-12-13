package lock_redis_set

import (
	"framework/class/lock"
	"framework/class/logger"
	"github.com/go-redis/redis/v8"
)

type Options struct {
	redisClient *redis.Client
	logger      logger.Logger
}

type Option func(o *Options)

func NewLock(key string, val string,opts ...Option) lock.Lock {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	p := &plugin{
		key:         key,
		val:         val,
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

func WithLogger(log logger.Logger) Option {
	return func(o *Options) {
		o.logger = log
	}
}
