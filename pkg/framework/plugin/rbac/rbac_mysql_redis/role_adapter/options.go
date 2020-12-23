package role_adapter

import (
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
)

type Options struct {
	roleModel               rbac.Model
	roleAssocApiModel       rbac.Model
	roleAssocMenuGroupModel rbac.Model
	roleAssocMenuModel      rbac.Model
	dbMaster                *gorm.DB
	dbSlave                 *gorm.DB
	*Callback
}

type Option func(o *Options)

func New(opts ...Option) rbac.RoleAdapter {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	p := &plugin{
		roleModel:               nil,
		roleAssocApiModel:       nil,
		roleAssocMenuGroupModel: nil,
		roleAssocMenuModel:      nil,
		Callback:                nil,
		dbMaster:                nil,
		dbSlave:                 nil,
		opts:                    options,
	}
	if err := p.Init(); err != nil {
		panic(err)
	}
	return p
}

func WithModel(role rbac.Model, roleAssocApi rbac.Model, roleAssocMenuGroup rbac.Model, roleAssocMenu rbac.Model) Option {
	return func(o *Options) {
		o.roleModel = role
		o.roleAssocApiModel = roleAssocApi
		o.roleAssocMenuGroupModel = roleAssocMenuGroup
		o.roleAssocMenuModel = roleAssocMenu
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
