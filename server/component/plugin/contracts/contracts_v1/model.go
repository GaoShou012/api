package contracts_v1

import (
	"gorm.io/gorm"
	"time"
)

// 客户模型
type CustomerModel interface {
	GetExpirationDateWithLock(tx *gorm.DB) (time.Time, error)
	SetExpirationDateWithTx(tx *gorm.DB, expiration time.Time) error
}

// 合同模型
type ContractModel interface {
	GetTableName() string
	GetContractLife() (time.Time, time.Time)
	Invalid(id uint64, args ...interface{}) error
	Review(id uint64, args ...interface{}) error
	EffectiveWithTx(tx *gorm.DB, id uint64) error
	TerminateWithTx(tx *gorm.DB, id uint64) error
}
