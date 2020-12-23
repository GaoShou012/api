package models

import "time"

type RbacRoleAssocApi struct {
	Id        *uint64
	TenantId  *uint64
	RoleId    *uint64
	ApiId     *uint64
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (m *RbacRoleAssocMenu) GetTableName() string {
	return "rbac_role_assoc_api"
}
