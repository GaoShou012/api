package models_tenant

import "api/models"

type RbacApi struct {
	models.RbacApi
	TenantId *uint64
}

func (m *RbacApi) GetTableName() string {
	return "tenants_rbac_api"
}

func (m *RbacApi) GetId() uint64 {
	return *m.Id
}
func (m *RbacApi) GetMethod() string {
	return *m.Method
}
func (m *RbacApi) GetPath() string {
	return *m.Path
}

func (m *RbacApi) BeforeUpdate() {
	m.TenantId = nil
	m.Model.BeforeUpdate()
}
