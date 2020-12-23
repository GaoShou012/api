package menu_adapter

import (
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
)

type Options struct {
	menuModel      rbac.Menu
	menuGroupModel rbac.MenuGroup
	dbMaster       *gorm.DB
	dbSlave        *gorm.DB
	*Callback
}

type Option func(o *Options)

func New(opts ...Option) rbac.MenuAdapter {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	p := &plugin{
		menuModel:      nil,
		menuGroupModel: nil,
		dbMaster:       nil,
		dbSlave:        nil,
		Callback:       nil,
		opts:           options,
	}
	if err := p.Init(); err != nil {
		panic(err)
	}
	return p
}

func WithModel(menuModel rbac.Menu, menuGroupModel rbac.MenuGroup) Option {
	return func(o *Options) {
		o.menuModel = menuModel
		o.menuGroupModel = menuGroupModel
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
