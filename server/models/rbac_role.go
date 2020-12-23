package models

import "time"

type RbacRole struct {
	Id        *uint64
	TenantId  *uint64
	Name      *string
	Desc      *string
	Icon      *string
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

func (m *RbacRole) GetTableName() string {
	return "rbac_role"
}
