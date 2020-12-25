package middleware_gin

import (
	"framework/class/middleware"
	"github.com/go-redis/redis/v8"
)

type Options struct {
	model          middleware.Operator
	redisClient    *redis.Client
	cipherKey      []byte
	headerTokenKey string
	*Callback
}

type Option func(o *Options)

func New(opts ...Option) middleware.OperatorContext {
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

// 解析Token时使用的密码本
func WithCipherKey(key []byte) Option {
	return func(o *Options) {
		o.cipherKey = key
	}
}

// Token，存放在上下文的key
func WithHeaderTokenKey(key string) Option {
	return func(o *Options) {
		o.headerTokenKey = key
	}
}

// 操作者 数据模型
func WithModel(model middleware.Operator) Option {
	return func(o *Options) {
		o.model = model
	}
}

// 分布式使用的超时检测
func WithRedisClient(redisClient *redis.Client) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}

// 回调
func WithCallback(callback *Callback) Option {
	return func(o *Options) {
		o.Callback = callback
	}
}
