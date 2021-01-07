package models_merchant

import "api/models"

type RbacMenu struct {
	models.RbacMenu
	TenantId *uint64
}

func (m *RbacMenu) GetTableName() string {
	return "tenants_rbac_menu"
}

func (m *RbacMenu) BeforeUpdate() {
	m.TenantId = nil
	m.Model.BeforeUpdate()
}
