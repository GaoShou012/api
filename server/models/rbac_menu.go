package models

type RbacMenu struct {
	Model
	TenantId *uint64
	GroupId  *uint64
	Sort     *uint64
	Name     *string
	Code     *string
	Icon     *string
	Desc     *string
}

func (m *RbacMenu) GetTableName() string {
	return "rbac_menu"
}

func (m *RbacMenu) BeforeUpdate() {
	m.TenantId = nil
	m.Model.BeforeUpdate()
}
