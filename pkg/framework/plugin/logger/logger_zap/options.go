package logger_zap

import (
	"framework/class/logger"
	"go.uber.org/zap"
)

type Options struct {
	*zap.SugaredLogger
}

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

func WithSugraedLogger(loggerClient *zap.SugaredLogger) Option {
	return func(o *Options) {
		o.SugaredLogger = loggerClient
	}
}
