package models

import (
	"api/global"
	"time"
)

type Merchants struct {
	Model
	Name       *string
	Code       *string
	Expiration *time.Time
	Channel    *int
	Enable     *bool
	MaxVisitor *int
	Desc       *string
}

func (m *Merchants) GetTableName() string {
	return "merchants"
}

func (m *Merchants) SelectByCode(fields string, code string) (bool, error) {
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

// 商户是否启用
func (m *Merchants) IsEnable() bool {
	return *m.Enable
}

// 商户是否租约过期
func (m *Merchants) IsExpiration() bool {
	if m.Expiration.Before(time.Now()) {
		return false
	}
	return true
}
