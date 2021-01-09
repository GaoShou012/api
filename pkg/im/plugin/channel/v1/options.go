package channel_v1

import (
	"framework/class/sortedset"
	"framework/class/stream"
	"github.com/go-redis/redis"
	"im/class/channel"
)

type Options struct {
	//infoModel   channel.Info
	redisClient *redis.Client
	sortedset.Sortedset
	stream.Stream
}

type Option func(o *Options)

func New(opts ...Option) channel.Channel {
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

func WithStream(stream stream.Stream) Option {
	return func(o *Options) {
		o.Stream = stream
	}
}
