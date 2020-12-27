package models

type RbacMenuGroup struct {
	Model
	Sort     *uint64
	Name     *string
	Code    *string
	Icon     *string
	Desc     *string
}

func (m *RbacMenuGroup) GetTableName() string {
	return "rbac_menu_group"
}

