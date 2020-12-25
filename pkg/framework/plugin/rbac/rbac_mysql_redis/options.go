package rbac_mysql_redis

import (
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
)

type Options struct {
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
}

type Option func(o *Options)

func New(opts ...Option) rbac.RBAC {
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

func WithGorm(dbMaster *gorm.DB, dbSlave *gorm.DB) Option {
	return func(o *Options) {
		o.dbMaster = dbMaster
		o.dbSlave = dbSlave
	}
}
