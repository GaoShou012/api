package models

import (
	"api/utils"
	"time"
)

type AuthorityRolesApis struct {
	Id        *uint64
	RoleId    *uint64 // 菜单名称
	ApiMethod *string // 排序值
	ApiPath   *string // 菜单图标
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
func (m *AuthorityRolesApis) UpdateById(param *AuthorityRolesApis) error {
	res := utils.IMysql.Master.Model(m).Updates(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *AuthorityRolesApis) DeleteById(param *AuthorityRolesApis) error {
	res := utils.IMysql.Master.Model(m).Delete(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *AuthorityRolesApis) Count(field string) (int, error) {
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