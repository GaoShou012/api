package contracts_v1

import (
	"api/component/class/contracts"
	"fmt"
	"time"
)

var _ contracts.Contracts = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Create(contract contracts.Contract) error {
	tableName := p.opts.contractModel.GetTableName()
	res := p.opts.dbMaster.Table(tableName).Create(contract)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *plugin) Review(id uint64, args ...interface{}) error {
	return p.opts.contractModel.Review(id, args...)
}

func (p *plugin) Effective(contractId uint64, args ...interface{}) error {
	now := time.Now()
	effective, expiration := p.opts.contractModel.GetContractLife()

	tx := p.opts.dbMaster.Begin()
	expirationOfCustomer, err := p.opts.customerModel.GetExpirationDateWithLock(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 叠加有效时间
	{
		if effective.After(expiration) {
			err := fmt.Errorf("生效日期(%s) 不能大于 过期日期(%s)", effective, expiration)
			return err
		}
		if expiration.Before(now) {
			err := fmt.Errorf("过期日期(%s) 不能少于 当前时间(%s)", expiration, now)
			return err
		}
		var _duration time.Duration
		if now.After(effective) {
			_duration = expiration.Sub(now)
		} else {
			_duration = expiration.Sub(effective)
		}
		expirationOfCustomer.Add(_duration)
	}

	// 保存客户截止日期
	if err := p.opts.customerModel.SetExpirationDateWithTx(tx, expirationOfCustomer); err != nil {
		tx.Rollback()
		return err
	}

	// 合同生效
	if err := p.opts.contractModel.EffectiveWithTx(tx, contractId); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (p *plugin) Invalid(contractId uint64, args ...interface{}) error {
	return p.opts.contractModel.Invalid(contractId, args...)
}

func (p *plugin) Terminate(contractId uint64, args ...interface{}) error {
	now := time.Now()
	_, expiration := p.opts.contractModel.GetContractLife()

	tx := p.opts.dbMaster.Begin()
	expirationOfCustomer, err := p.opts.customerModel.GetExpirationDateWithLock(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 减少有效时间
	{
		if expiration.Before(now) {
			err := fmt.Errorf("合同已经过期，不能进行终止")
			return err
		}

		var _duration time.Duration
		_duration = expirationOfCustomer.Sub(expiration)
		expirationOfCustomer = expirationOfCustomer.Add(_duration)
	}

	// 保存客户截止日期
	if err := p.opts.customerModel.SetExpirationDateWithTx(tx, expirationOfCustomer); err != nil {
		tx.Rollback()
		return err
	}

	// 合同终止
	if err := p.opts.contractModel.TerminateWithTx(tx, contractId); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
