package models

type RbacApi struct {
	Model
	TenantId *uint64
	Method   *string
	Path     *string
}

func (m *RbacApi) GetTableName() string {
	return "rbac_api"
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
