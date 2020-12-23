package api_adapter

import (
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
)

type Options struct {
	model    rbac.Api
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
	*Callback
}

type Option func(o *Options)

func New(opts ...Option) rbac.ApiAdapter {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	p := &plugin{
		model:    nil,
		dbMaster: nil,
		dbSlave:  nil,
		Callback: nil,
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

func WithCallback(callback *Callback) Option {
	return func(o *Options) {
		o.Callback = callback
	}
}
