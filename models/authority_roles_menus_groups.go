package models

import (
	"api/utils"
	"time"
)

type AuthorityRolesMenusGroups struct {
	Id         *uint64
	RoleId     *uint64
	MenuGroups *string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
func (m *AuthorityRolesMenusGroups) UpdateById(param *AuthorityRolesMenusGroups) error {
	res := utils.IMysql.Master.Model(m).Updates(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *AuthorityRolesMenusGroups) DeleteById(param *AuthorityRolesMenusGroups) error {
	res := utils.IMysql.Master.Model(m).Delete(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *AuthorityRolesMenusGroups) Count(field string) (int, error) {
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