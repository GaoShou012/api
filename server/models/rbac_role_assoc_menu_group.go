package models

type RbacRoleAssocMenuGroup struct {
	Model
	RoleId      *uint64
	MenuGroupId *uint64
}

func (m *RbacRoleAssocMenuGroup) GetTableName() string {
	return "rbac_role_assoc_menu_group"
}
