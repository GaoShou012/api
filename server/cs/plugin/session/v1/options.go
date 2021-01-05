package v1

import (
	"api/cs/class/session"
	"framework/class/stream"
)

type Options struct {
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

func WithStream(stream stream.Stream) Option {
	return func(o *Options) {
		o.Stream = stream
	}
}
