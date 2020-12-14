package models

import (
	"api/global"
	"time"
)

type MenuGroups struct {
	Id        *uint64
	Group     *string
	Sort      *int
	Icon      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
func (m *MenuGroups) UpdateById(param *MenuGroups) error {
	res := global.DBMaster.Model(m).Updates(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (m *MenuGroups) DeleteById(param *MenuGroups) error {
	res := global.DBMaster.Model(m).Delete(param)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *MenuGroups) Count(field string) (int64, error) {
	count := 0
	res :=global.DBSlave.Model(m)
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
