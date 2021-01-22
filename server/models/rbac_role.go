package models

import "api/global"

type RbacRole struct {
	Model
	Name   *string
	Enable *bool
	Desc   *string
	Icon   *string
}

func (m *RbacRole) GetTableName() string {
	return "rbac_role"
}

func (m *RbacRole) SelectByName(fields string, name string) error {
	res := global.DBMaster.Table(m.GetTableName()).Select(fields).Where("name=?", name).First(m)
	return res.Error
}

func (m *RbacRole) Insert() error {
	res := global.DBMaster.Table(m.GetTableName()).Create(m)
	return res.Error
}
