package channel_v1

import (
	"framework/class/channel"
	"framework/class/stream"
	"github.com/go-redis/redis/v8"
)

type Options struct {
	channel.ClientAdapter
	redisClient *redis.Client
	*Callback
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
	}
}

func WithStream(stream stream.Stream) Option {
	return func(o *Options) {
		o.Stream = stream
	}
}

func WithCallback(callback *Callback) Option {
	return func(o *Options) {
		o.Callback = callback
	}
}

//func WithLogger(log logger.Logger) Option {
//	return func(o *Options) {
//		o.logger = log
//	}
//}
