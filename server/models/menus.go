package models

import (
	"api/global"
	"time"
)

// Menu 菜单实体
type Menus struct {
	Id        *uint64
	Name      *string
	Icon      *string
	Sort      *int
	Path      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (m *Menus) UpdateById(param *Menus) error {
	res := global.DBMaster.Model(m).Updates(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *Menus) DeleteById(param *Menus) error {
	res := global.DBMaster.Model(m).Delete(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *Menus) Count(field string) (int, error) {
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
	return count, nil
}