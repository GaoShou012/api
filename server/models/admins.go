package models

import (
	"api/global"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	AdminsUserTypeSuperManager = iota
	AdminsUserTypeManager
)

type Admins struct {
	Id        *uint64    // 自增ID
	Enable    *bool      `gorm:"default:false"` // 是否启用 0=不启用，1=启用
	Username  *string    // 账号
	Password  *string    // 密码
	Nickname  *string    // 昵称
	Roles     *string    // 关联的角色ID，多角色使用逗号间隔
	CreatedAt *time.Time // 创建时间
	UpdatedAt *time.Time // 更新时间
}

func (m *Admins) GetTableName() string {
	return "admins"
}

// 密码加密
func (m *Admins) EncryptPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("加密错误")
	}
	encryptPassword := string(hashPassword)
	return encryptPassword, nil
}

// 账号是否启用
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

// 账号是否存在
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

// 根据ID，查询数据
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

// 根据账号，查询数据
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
