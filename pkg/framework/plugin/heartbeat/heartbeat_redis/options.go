package heartbeat_redis

import (
	"framework/class/heartbeat"
	"github.com/go-redis/redis"
	"time"
)

type Option func(o *Options)
type Options struct {
	timeout     time.Duration
	redisClient *redis.Client
}

func New(opts ...Option) heartbeat.Heartbeat {
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

func WithRedisClient(redisClient *redis.Client) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.timeout = timeout
	}
}

//func WithLogger(log logger.Logger) Option {
//	return func(o *Options) {
//		o.logger = log
//	}
//}
