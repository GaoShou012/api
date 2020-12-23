package models

import "time"

type RbacRoleAssocMenu struct {
	Id        *uint64
	TenantId  *uint64
	RoleId    *uint64
	MenuId    *uint64
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (m *RbacRoleAssocApi) GetTableName() string {
	return "rbac_role_assoc_menu"
}
