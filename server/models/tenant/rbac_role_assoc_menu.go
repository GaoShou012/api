package models_tenant

import "api/models"

type RbacRoleAssocMenu struct {
	models.RbacRoleAssocMenu
}

func (m *RbacRoleAssocMenu) GetTableName() string {
	return "tenants_rbac_role_assoc_menu"
}

