package countdown_context

import (
	"framework/class/countdown"
	"github.com/go-redis/redis"
	"time"
)

type Option func(o *Options)
type Options struct {
	timeout     time.Duration
	redisClient *redis.Client
}

func New(opts ...Option) countdown.Countdown {
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
