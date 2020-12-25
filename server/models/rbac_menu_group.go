package models

type RbacMenuGroup struct {
	Model
	TenantId *uint64
	Sort     *uint64
	Name     *string
	Icon     *string
	Desc     *string
}

func (m *RbacMenuGroup) GetTableName() string {
	return "rbac_menu_group"
}

func (m *RbacMenuGroup) BeforeUpdate() {
	m.TenantId = nil
	m.Model.BeforeUpdate()
}
