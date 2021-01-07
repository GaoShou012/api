package models_merchant

import "api/models"

type RbacMenuGroup struct {
	models.RbacMenuGroup
	TenantId *uint64
}

func (m *RbacMenuGroup) GetTableName() string {
	return "tenants_rbac_menu_group"
}

func (m *RbacMenuGroup) BeforeUpdate() {
	m.TenantId = nil
	m.Model.BeforeUpdate()
}
