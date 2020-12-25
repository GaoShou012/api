package models

type RbacRoleAssocMenu struct {
	Model
	TenantId *uint64
	RoleId   *uint64
	MenuId   *uint64
}

func (m *RbacRoleAssocMenu) GetTableName() string {
	return "rbac_role_assoc_menu"
}

func (m *RbacRoleAssocMenu) BeforeUpdate() {
	m.TenantId = nil
	m.Model.BeforeUpdate()
}
