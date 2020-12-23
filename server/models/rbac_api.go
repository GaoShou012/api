package models

import "time"

type RbacApi struct {
	Id        *uint64
	TenantId  *uint64
	Method    *string
	Path      *string
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

func (m *RbacApi) GetTableName() string {
	return "rbac_api"
}
