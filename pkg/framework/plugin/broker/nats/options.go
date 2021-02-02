package broker_nats

import (
	"framework/class/broker"
	"github.com/nats-io/nats.go"
)

type Options struct {
	natsClient *nats.Conn
}

type Option func(o *Options)

func New(opts ...Option) broker.Broker {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	b := &plugin{
		opts: options,
	}
	if err := b.Init(); err != nil {
		panic(err)
	}
	return b
}

func WithNatsClient(natsClient *nats.Conn) Option {
	return func(o *Options) {
		o.natsClient = natsClient
	}
}
