package session_redis

import (
	"cs/class/session"
	"im/class/channel"
	"framework/class/stream"
	"github.com/go-redis/redis/v8"
)

type Options struct {
	channel.Channel
	redisClient *redis.Client
	stream.Stream
}

type Option func(o *Options)

func New(opts ...Option) session.Session {
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

//func WithLogger(log logger.Logger) Option {
//	return func(o *Options) {
//		o.logger = log
//	}
//}
