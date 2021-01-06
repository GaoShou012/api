package im_v1

import (
	"framework/class/broker"
	"github.com/go-redis/redis"
	"im/class/channel"
	"im/class/client"
	"im/class/gateway"
	"im/class/im"
)

type Options struct {
	// 用于分发消息到网关
	broker broker.Broker
	// 客户端消息
	client      client.Client
	channel     channel.Channel
	gateway     gateway.Gateway
	redisClient *redis.Client
}

type Option func(o *Options)

func New(opts ...Option) im.IM {
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

func WithBroker(broker2 broker.Broker) Option {
	return func(o *Options) {
		o.broker = broker2
	}
}

func WithChannel(channel2 channel.Channel) Option {
	return func(o *Options) {
		o.channel = channel2
	}
}

func WithClient(client2 client.Client) Option {
	return func(o *Options) {
		o.client = client2
	}
}

func WithGateway(gateway2 gateway.Gateway) Option {
	return func(o *Options) {
		o.gateway = gateway2
	}
}

func WithRedisClient(client2 *redis.Client) Option {
	return func(o *Options) {
		o.redisClient = client2
	}
}
