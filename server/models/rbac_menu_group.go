package models

import "time"

type RbacMenuGroup struct {
	Id        *uint64
	TenantId  *uint64
	Sort      *uint64
	Name      *string
	Icon      *string
	Desc      *string
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

func (m *RbacMenuGroup) GetTableName() string {
	return "rbac_menu_group"
}
