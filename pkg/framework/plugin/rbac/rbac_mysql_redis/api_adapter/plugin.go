package api_adapter

import (
	"errors"
	"fmt"
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
	"reflect"
)

var _ rbac.ApiAdapter = &plugin{}

type plugin struct {
	model rbac.Api
	*Callback
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
	opts     *Options
}

func (p *plugin) Init() error {
	if p.opts.model == nil {
		return errors.New("model is nil")
	}
	p.model = p.opts.model
	if p.opts.Callback == nil {
		return errors.New("callback is nil")
	}
	p.Callback = p.opts.Callback
	if p.opts.dbMaster == nil {
		return errors.New("db master is nil")
	}
	p.dbMaster = p.opts.dbMaster
	if p.opts.dbSlave == nil {
		return errors.New("db slave is nil")
	}
	p.dbSlave = p.opts.dbSlave
	return nil
}

func (p *plugin) Authority(operator rbac.Operator, apiId uint64) (bool, error) {
	return p.Callback.Authority(operator, apiId)
}

func (p *plugin) Create(api rbac.Api) error {
	res := p.dbMaster.Table(p.model.GetTableName()).Create(api)
	return res.Error
}

func (p *plugin) Delete(apiId uint64) error {
	res := p.dbMaster.Table(p.model.GetTableName()).Where("id=?", apiId).Delete(p.model)
	return res.Error
}

func (p *plugin) Update(apiId uint64, api rbac.Api) error {
	res := p.dbMaster.Table(p.model.GetTableName()).Where("id=?", apiId).Updates(api)
	if res.RowsAffected != 1 {
		return fmt.Errorf("更新失败，ID(%d)不存在", apiId)
	}
	return res.Error
}

func (p *plugin) SelectById(apiId uint64) (rbac.Api, error) {
	t := reflect.TypeOf(p.model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	api := reflect.New(t).Interface()
	res := p.dbSlave.Table(p.model.GetTableName()).Where("id=?", apiId).Find(api)
	if res.Error != nil {
		return nil, res.Error
	}
	return api.(rbac.Api), nil
}

func (p *plugin) FindById(operator rbac.Operator, apiId uint64, api rbac.Api) error {
	res := p.dbSlave.Table(p.model.GetTableName()).Where("id=?", apiId).Find(api)
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
