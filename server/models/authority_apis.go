package models

import (
	"api/utils"
	"time"
)

type AuthorityApis struct {
	Id        *uint64
	Method    *string
	Path      *string // 访问路由
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
func (m *AuthorityApis) UpdateById(param *AuthorityApis) error {
	res := utils.IMysql.Master.Model(m).Updates(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *AuthorityApis) DeleteById(param *AuthorityApis) error {
	res := utils.IMysql.Master.Model(m).Delete(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *AuthorityApis) Count(field string) (int, error) {
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