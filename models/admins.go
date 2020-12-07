package models

import (
	"api/utils"
	"fmt"
	"time"
)

type Admins struct {
	Id        *uint64
	Enable    *bool
	State     *uint64
	UserType  *uint64
	Username  *string
	Password  *string
	Nickname  *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (m *Admins) IsEnable(id uint64) (bool, error) {
	admin := &Admins{}
	res := utils.IMysql.Slave.Select("enable").Where("id=?", id).First(admin)
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
	res := utils.IMysql.Slave.Model(m).Where("username=?", username).Count(&count)
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
	res := utils.IMysql.Slave.Select(fields).Where("id=?", id).First(m)
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
	res := utils.IMysql.Slave.Select(fields).Where("username=?", username).First(m)
	if res.Error != nil {
		if res.RecordNotFound() {
			return fmt.Errorf("账号不存在")
		} else {
			return res.Error
		}
	}
	return nil
}
