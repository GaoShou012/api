package logger_log

import (
	"framework/class/logger"
)

type Options struct{}

type Option func(o *Options)

func New(opts ...Option) logger.Logger {
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
