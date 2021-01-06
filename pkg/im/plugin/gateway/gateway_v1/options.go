package gateway_v1

import (
	"framework/class/broker"
	"im/class/gateway"
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

// 如果初始化带topic
// gateway使用此topic进行pub sub
// 如果不带topic，gateway使用 topic="gateway" 进行 pub sub
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
