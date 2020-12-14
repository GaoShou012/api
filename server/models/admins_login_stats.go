package models

import (
	"api/global"
	"github.com/jinzhu/gorm"
	"time"
)

type AdminsLoginStats struct {
	Id *uint64

	// 用户ID
	UserId *uint64

	// 登陆次数
	LoginTimes *uint64 `gorm:"default:0"`

	UpdatedAt *time.Time
	CreatedAt *time.Time
}

func (m *AdminsLoginStats) Create(userId uint64) *gorm.DB {
	i := &AdminsLoginStats{
		Id:         nil,
		UserId:     &userId,
		LoginTimes: nil,
		UpdatedAt:  nil,
		CreatedAt:  nil,
	}
	return global.DBMaster.Create(i)
}

func (m *AdminsLoginStats) IsExistsByUserId(userId uint64) (bool, error) {
	count := 0
	res := global.DBSlave.Model(m).Where("user_id=?", userId).Count(&count)
	if res.Error != nil {
		return false, res.Error
	}
	if count == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (m *AdminsLoginStats) IncrByUserId(id uint64) *gorm.DB {
	return global.DBMaster.Exec("update `admins_login_stats` set login_times=login_times+1 where user_id=?", id)
}
