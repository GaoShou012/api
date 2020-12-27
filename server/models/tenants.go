package models

import (
	"api/global"
	"time"
)

type Tenants struct {
	Model
	Enable     *bool      // 是否启用
	Expiration *time.Time // 租户的截止时间，可以多合同叠加，默认是null，表示还没有生效
	Code       *string    // 租户编码
	Name       *string    // 租户名称
	Desc       *string    // 描述
}

func (m *Tenants) GetTableName() string {
	return "tenants"
}

func (m *Tenants) SelectByCode(fields string, code string) (bool, error) {
	res := global.DBSlave.Table(m.GetTableName()).Select(fields).Where("code=?", code).First(m)
	if res.Error != nil {
		if res.RecordNotFound() {
			return false, nil
		} else {
			return false, res.Error
		}
	}
	return true, nil
}
