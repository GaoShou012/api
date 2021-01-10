package models_merchant

import "api/models"

type RbacRoleAssocMenuGroup struct {
	models.RbacRoleAssocMenuGroup
}

func (m *RbacRoleAssocMenuGroup) GetTableName() string {
	return "merchants_rbac_role_assoc_menu_group"
}
