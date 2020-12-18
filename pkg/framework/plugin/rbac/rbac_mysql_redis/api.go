package rbac_mysql_redis

import (
	"framework/class/rbac"
	"github.com/jinzhu/gorm"
)

var _ rbac.ApiAdapter = &ApiAdapter{}

type ApiAdapter struct {
	db   *gorm.DB
	opts *Options
}

func (p *ApiAdapter) Init() error {
	panic("implement me")
}

func (p *ApiAdapter) Create(api rbac.Api) error {
	res := p.db.Model(api).Create(api)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *ApiAdapter) Update(apiId uint64, api rbac.Api) error {
	res := p.db.Model(api).Where("id=?", apiId).Updates(api)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *ApiAdapter) SelectById(tenantId uint64, apiId uint64) (rbac.Api, error) {
	panic("implement me")
}

func (p *ApiAdapter) SelectByPage(tenantId uint64, page uint64, pageSize uint64) ([]rbac.Api, error) {
	panic("implement me")
}

func (p *ApiAdapter) Count(tenantId uint64) (uint64, error) {
	panic("implement me")
}

func (p *ApiAdapter) VerifyIdWithOperator(id uint64, operator rbac.Operator) (bool, error) {
	panic("implement me")
}
