package models_tenant

import "api/models"

type RbacRoleAssocApi struct {
	models.RbacRoleAssocApi
}

func (m *RbacRoleAssocApi) GetTableName() string {
	return "tenants_rbac_role_assoc_api"
}

func (m *RbacRoleAssocApi) GetRoleId() uint64 {
	return *m.RoleId
}
func (m *RbacRoleAssocApi) GetApiId() uint64 {
	return *m.ApiId
}