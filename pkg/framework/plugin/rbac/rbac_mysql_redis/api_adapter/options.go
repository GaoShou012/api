package api_adapter

import (
	"framework/class/rbac"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Options struct {
	model       rbac.Api
	dbMaster    *gorm.DB
	dbSlave     *gorm.DB
	redisClient *redis.Client
	*Callback
}

type Option func(o *Options)

func New(opts ...Option) rbac.ApiAdapter {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	p := &plugin{
		opts:     options,
	}
	if err := p.Init(); err != nil {
		panic(err)
	}
	return p
}

func WithModel(model rbac.Api) Option {
	return func(o *Options) {
		o.model = model
	}
}

func WithGorm(master *gorm.DB, slave *gorm.DB) Option {
	return func(o *Options) {
		o.dbMaster = master
		o.dbSlave = slave
	}
}

func WithRedisClient(redisClient *redis.Client) Option {
	return func(o *Options) {
		o.redisClient = redisClient
	}
}

func WithCallback(callback *Callback) Option {
	return func(o *Options) {
		o.Callback = callback
	}
}
