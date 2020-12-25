package models

type RbacRole struct {
	Model
	TenantId *uint64
	Name     *string
	Desc     *string
	Icon     *string
}

func (m *RbacRole) GetTableName() string {
	return "rbac_role"
}

func (m *RbacRole) BeforeUpdate() {
	m.TenantId = nil
	m.Model.BeforeUpdate()
}
