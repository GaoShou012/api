package im_v1

import (
	"github.com/go-redis/redis"
	"im/class/channel"
	"im/class/client"
	"im/class/gateway"
	"im/class/im"
)

type Options struct {
	clientAdapter  client.Client
	client         client.Client
	channelAdapter channel.Channel
	gateway        gateway.Gateway
	redisClient    *redis.Client
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

func WithClientAdapter(clientAdapter client.Client) Option {
	return func(o *Options) {
		o.clientAdapter = clientAdapter
	}
}

func WithChannelAdapter(channelAdapter channel.Channel) Option {
	return func(o *Options) {
		o.channelAdapter = channelAdapter
	}
}
