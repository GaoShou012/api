package models_merchant

import (
	"api/global"
	"api/models"
)

type Admins struct {
	models.Model
	Enable   *bool   `gorm:"default:false"` // 是否启用 0=不启用，1=启用
	State    *uint64 `gorm:"default:0"`     // 状态 0=未初始化，1=正常，2=冻结
	UserType *uint64 // 用户类型 根据项目类型进行定义
	Username *string // 账号
	Password *string // 密码
	Nickname *string // 昵称
}

func (m *Admins) GetTableName() string {
	return "merchants_admins"
}

func (m *Admins) SelectByUsername(fields string, tenantId uint64, username string) (bool, error) {
	res := global.DBSlave.Select(fields).Where("merchant_id =? and username=?", tenantId, username).First(m)
	if res.Error != nil {
		if res.RecordNotFound() {
			return false, nil
		} else {
			return false, res.Error
		}
	}
	return true, nil
}
