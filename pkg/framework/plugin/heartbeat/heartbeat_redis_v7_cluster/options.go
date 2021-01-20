package heartbeat_redis_v7_cluster

import (
	"github.com/go-redis/redis/v7"
	"framework/class/heartbeat"
	"time"
)

type Option func(o *Options)
type Options struct {
	timeout     time.Duration
	redisClient *redis.ClusterClient
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

func WithRedisClient(redisClient *redis.ClusterClient) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.timeout = timeout
	}
}

