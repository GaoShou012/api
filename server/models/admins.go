package models

import (
	"api/global"
	"fmt"
	"time"
)

type Admins struct {
	Id        *uint64 // 自增ID
	Enable    *bool   `gorm:"default:false"` // 是否启用 0=不启用，1=启用
	State     *uint64 `gorm:"default:0"`     // 状态 0=未初始化，1=正常，2=冻结
	UserType  *uint64 // 用户类型 根据项目类型进行定义
	Role      *string
	Username  *string    // 账号
	Password  *string    // 密码
	Nickname  *string    // 昵称
	CreatedAt *time.Time // 创建时间
	UpdatedAt *time.Time // 更新时间
}

func (m *Admins) IsEnable(id uint64) (bool, error) {
	admin := &Admins{}
	res := global.DBSlave.Select("enable").Where("id=?", id).First(admin)
	if res.Error != nil {
		if res.RecordNotFound() {
			return false, fmt.Errorf("ID(%d)不存在", id)
		} else {
			return false, res.Error
		}
	}
	return *admin.Enable, nil
}

func (m *Admins) IsExistsByUsername(username string) (bool, error) {
	count := 0
	res := global.DBSlave.Model(m).Where("username=?", username).Count(&count)
	if res.Error != nil {
		return false, res.Error
	}
	if count == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (m *Admins) SelectById(fields string, id uint64) error {
	res := global.DBSlave.Select(fields).Where("id=?", id).First(m)
	if res.Error != nil {
		if res.RecordNotFound() {
			return fmt.Errorf("ID(%d)不存在", id)
		} else {
			return res.Error
		}
	}
	return nil
}

func (m *Admins) SelectByUsername(fields string, username string) error {
	res := global.DBSlave.Select(fields).Where("username=?", username).First(m)
	if res.Error != nil {
		if res.RecordNotFound() {
			return fmt.Errorf("账号不存在")
		} else {
			return res.Error
		}
	}
	return nil
}
