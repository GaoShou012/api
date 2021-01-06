package client_v1

import (
	"framework/class/stream"
	"github.com/go-redis/redis"
	"im/class/client"
)

type Options struct {
	redisClient *redis.Client
	stream.Stream
}

type Option func(o *Options)

func New(opts ...Option) client.Client {
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
	}
}

func WithStream(stream stream.Stream) Option {
	return func(o *Options) {
		o.Stream = stream
	}
}
