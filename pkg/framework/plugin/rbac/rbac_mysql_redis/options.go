package rbac_mysql_redis

import (
	"framework/class/logger"
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
)

type Options struct {
	db     *gorm.DB
	logger logger.Logger
}

type Option func(o *Options)

func NewApiAdapter(opts ...Option) rbac.ApiAdapter {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	b := &ApiAdapter{
		db:   nil,
		opts: options,
	}
	if err := b.Init(); err != nil {
		panic(err)
	}
	return b
}

func WithDatabase(db *gorm.DB) Option {
	return func(o *Options) {
		o.db = db
	}
}

func WithLogger(log logger.Logger) Option {
	return func(o *Options) {
		o.logger = log
	}
}
