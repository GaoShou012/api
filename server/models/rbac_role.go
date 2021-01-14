package models

import "api/global"

type RbacRole struct {
	Model
	Name *string
	Desc *string
	Icon *string
}

func (m *RbacRole) GetTableName() string {
	return "rbac_role"
}

func (m *RbacRole) Insert() error {
	res := global.DBMaster.Table(m.GetTableName()).Create(m)
	return res.Error
}
