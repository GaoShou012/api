package models

import "api/global"

type RbacRoleAssocMenu struct {
	Model
	RoleId *uint64
	MenuId *uint64
}

func (m *RbacRoleAssocMenu) GetTableName() string {
	return "rbac_role_assoc_menu"
}

func (m *RbacRoleAssocMenu) Insert() error {
	res := global.DBMaster.Table(m.GetTableName()).Create(m)
	return res.Error
}
