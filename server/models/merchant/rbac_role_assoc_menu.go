package models_merchant

import "api/models"

type RbacRoleAssocMenu struct {
	models.RbacRoleAssocMenu
}

func (m *RbacRoleAssocMenu) GetTableName() string {
	return "merchants_rbac_role_assoc_menu"
}
