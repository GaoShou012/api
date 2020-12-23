package gateway_v1

import (
	"cs/class/gateway"
	"framework/class/broker"
)

type Options struct {
	topic string
	broker.Broker
}

type Option func(o *Options)

func New(opts ...Option) gateway.Gateway {
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

func WithTopic(topic string) Option {
	return func(o *Options) {
		o.topic = topic
	}
}

func WithBroker(b broker.Broker) Option {
	return func(o *Options) {
		o.Broker = b
	}
}
