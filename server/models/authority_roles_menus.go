package models

import (
	"api/global"
	"time"
)

type AuthorityRolesMenus struct {
	Id         *uint64
	RoleId     *uint64
	//MenuGroups *string
	MenuId     *uint64
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func (m *AuthorityRolesMenus) UpdateById(param *AuthorityRolesMenus) error {
	res := global.DBMaster.Model(m).Updates(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *AuthorityRolesMenus) DeleteById(param *AuthorityRolesMenus) error {
	res := global.DBMaster.Model(m).Delete(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *AuthorityRolesMenus) Count(field string) (int64, error) {
	count := 0
	res := global.DBSlave.Model(m)
	if field == "*" {
		res.Count(&count)
	} else {
		res.Where(field).Count(&count)
	}
	if res.Error != nil {
		return 0, res.Error
	}
	return int64(count), nil
}
