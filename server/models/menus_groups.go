package models

import (
	"api/global"
	"time"
)

type MenusGroups struct {
	Id        *uint64
	GroupName *string
	Sort      *uint64
	Icon      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}


func (m *MenusGroups) GetMenuName() string {
	return ""
}

func (m *MenusGroups) GetGroupName() string {
	return *m.GroupName
}
func (m *MenusGroups) GetSort() uint64 {
	return *m.Sort
}

func (m *MenusGroups) UpdateById(param *MenusGroups) error {
	res := global.DBMaster.Model(m).Updates(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (m *MenusGroups) DeleteById(param *MenusGroups) error {
	res := global.DBMaster.Model(m).Delete(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *MenusGroups) Count(field string) (int64, error) {
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
