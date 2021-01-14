package api_adapter

import (
	"errors"
	"fmt"
	"framework/class/rbac"
	lib_model "framework/libs/model"
)

var _ rbac.ApiAdapter = &plugin{}

type plugin struct {
	opts     *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Authority(operator rbac.Operator, apiId uint64) (bool, error) {
	return p.opts.Callback.Authority(operator, apiId)
}

func (p *plugin) Create(api rbac.Api) error {
	res := p.opts.dbMaster.Table(p.opts.model.GetTableName()).Create(api)
	return res.Error
}

func (p *plugin) Delete(apiId uint64) (bool, error) {
	res := p.opts.dbMaster.Table(p.opts.model.GetTableName()).Where("id=?", apiId).Delete(p.opts.model)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (p *plugin) Update(apiId uint64, api rbac.Api) error {
	api.BeforeUpdate()
	res := p.opts.dbMaster.Table(p.opts.model.GetTableName()).Where("id=?", apiId).Updates(api)
	if res.RowsAffected != 1 {
		return fmt.Errorf("更新失败，可能数据没有发生变化，或者ID(%d)不存在", apiId)
	}
	return res.Error
}

func (p *plugin) SelectById(apiId uint64) (rbac.Api, error) {
	newModel := lib_model.New(p.opts.model).(rbac.Api)
	res := p.opts.dbSlave.Table(p.opts.model.GetTableName()).Where("id=?", apiId).Find(newModel)
	if res.Error != nil {
		return nil, res.Error
	}
	return newModel, nil
}

func (p *plugin) FindById(operator rbac.Operator, apiId uint64, api rbac.Api) error {
	res := p.opts.dbSlave.Table(p.opts.model.GetTableName()).Where("id=?", apiId).Find(api)
	if res.Error != nil {
		return res.Error
	}
	ok, err := p.Authority(operator, apiId)
	if err != nil {
		return err
	}
	if !ok {
		api = nil
		return errors.New("权限不足")
	}
	return nil
}

func (p *plugin) FindByPage(operator rbac.Operator, page uint64, pageSize uint64, res []rbac.Api) error {
	panic("implement me")
}

func (p *plugin) Count(tenantId uint64) (uint64, error) {
	panic("implement me")
}
