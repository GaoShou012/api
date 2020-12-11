package models

import (
	"api/utils"
	"time"
)

// Menu 菜单实体
type Menus struct {
	Id        *uint64
	Name      *string // 菜单名称
	Sort      *int    // 排序值
	Icon      *string // 菜单图标
	Path      *string // 访问路由
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (m *Menus) UpdateById(menus *Menus) error {
	res := utils.IMysql.Master.Model(m).Updates(menus)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *Menus) Count(field string) (int, error) {
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
func (m *Menus) DeleteById(menus *Menus) error {
	res := utils.IMysql.Master.Model(m).Delete(menus)
	if res.Error != nil {
		return res.Error
	}
	return nil
}