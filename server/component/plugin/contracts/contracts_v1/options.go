package contracts_v1

import (
	"api/component/class/contracts"
	"framework/class/logger"
	"gorm.io/gorm"
)

type Options struct {
	customerModel CustomerModel
	contractModel ContractModel
	dbMaster      *gorm.DB
	dbSlave       *gorm.DB
	logger        logger.Logger
}

type Option func(o *Options)

func NewStream(opts ...Option) contracts.Contracts {
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

func WithModel(contractModel ContractModel,customerModel CustomerModel) Option {
	return func(o *Options) {
		o.contractModel = contractModel
		o.customerModel = customerModel
	}
}

func WithGorm(dbMaster *gorm.DB, dbSlave *gorm.DB) Option {
	return func(o *Options) {
		o.dbMaster = dbMaster
		o.dbSlave = dbSlave
	}
}
