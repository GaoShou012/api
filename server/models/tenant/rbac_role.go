package models_tenant

import "api/models"

type RbacRole struct {
	models.RbacRole
	TenantId *uint64
}

func (m *RbacRole) GetTableName() string {
	return "tenants_rbac_role"
}

func (m *RbacRole) BeforeUpdate() {
	m.TenantId = nil
	m.Model.BeforeUpdate()
}
