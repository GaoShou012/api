package models

import "time"

type RbacRoleAssocMenuGroup struct {
	Id          *uint64
	TenantId    *uint64
	RoleId      *uint64
	MenuGroupId *uint64
	UpdatedAt   *time.Time
	CreatedAt   *time.Time
}

func (m *RbacRoleAssocMenuGroup) GetTableName() string {
	return "rbac_role_assoc_menu_group"
}
