package models

type RbacApi struct {
	Model
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

