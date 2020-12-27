package models

type RbacRoleAssocMenu struct {
	Model
	RoleId   *uint64
	MenuId   *uint64
}

func (m *RbacRoleAssocMenu) GetTableName() string {
	return "rbac_role_assoc_menu"
}

