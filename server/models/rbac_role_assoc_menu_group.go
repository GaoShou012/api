package models

import "api/global"

type RbacRoleAssocMenuGroup struct {
	Model
	RoleId      *uint64
	MenuGroupId *uint64
}

func (m *RbacRoleAssocMenuGroup) GetTableName() string {
	return "rbac_role_assoc_menu_group"
}

func (m *RbacRoleAssocMenuGroup) Insert() error {
	res := global.DBMaster.Table(m.GetTableName()).Create(m)
	return res.Error
}
