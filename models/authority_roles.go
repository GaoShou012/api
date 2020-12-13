package models

import (
	"api/utils"
	"time"
)

type AuthorityRoles struct {
	Id        *uint64
	Name      *string
	Sort      *int    // 排序值
	Remark    *string // 备注
	Enable    *bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
func (m *AuthorityRoles) UpdateById(param *AuthorityRoles) error {
	res := utils.IMysql.Master.Model(m).Updates(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *AuthorityRoles) DeleteById(param *AuthorityRoles) error {
	res := utils.IMysql.Master.Model(m).Delete(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *AuthorityRoles) Count(field string) (int, error) {
	count := 0
	res := utils.IMysql.Slave.Model(m)
	if field == "*" {
		res.Count(&count)
	} else {
		res.Where(field).Count(&count)
	}
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}