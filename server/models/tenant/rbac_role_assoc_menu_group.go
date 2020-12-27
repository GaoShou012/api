package models_tenant

import "api/models"

type RbacRoleAssocMenuGroup struct {
	models.RbacRoleAssocMenuGroup
}

func (m *RbacRoleAssocMenuGroup) GetTableName() string {
	return "tenants_rbac_role_assoc_menu_group"
}
