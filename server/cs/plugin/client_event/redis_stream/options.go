package client_event_redis_stream

import (
	"api/cs/class/client_event"
	"framework/class/stream"
	"github.com/go-redis/redis/v8"
)

type Options struct {
	stream.Stream
}

type Option func(o *Options)

func New(opts ...Option) client_event.ClientEvent {
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
