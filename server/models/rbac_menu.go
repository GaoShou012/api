package models

import "time"

type RbacMenu struct {
	Id        *uint64
	TenantId  *uint64
	GroupId   *uint64
	Sort      *uint64
	Name      *string
	Icon      *string
	Desc      *string
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

func (m *RbacMenu) GetTableName() string {
	return "rbac_menu"
}
