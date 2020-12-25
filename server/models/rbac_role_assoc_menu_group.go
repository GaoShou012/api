package models

type RbacRoleAssocMenuGroup struct {
	Model
	TenantId    *uint64
	RoleId      *uint64
	MenuGroupId *uint64
}

func (m *RbacRoleAssocMenuGroup) GetTableName() string {
	return "rbac_role_assoc_menu_group"
}
func (m *RbacRoleAssocMenuGroup) BeforeUpdate() {
	m.TenantId = nil
	m.BeforeUpdate()
}
