package models

type RbacRoleAssocApi struct {
	Model
	//TenantId *uint64
	RoleId   *uint64
	ApiId    *uint64
}

func (m *RbacRoleAssocApi) GetTableName() string {
	return "rbac_role_assoc_api"
}

func (m *RbacRoleAssocApi) GetRoleId() uint64 {
	return *m.RoleId
}
func (m *RbacRoleAssocApi) GetApiId() uint64 {
	return *m.ApiId
}

func (m *RbacRoleAssocApi) BeforeUpdate() {
	//m.TenantId = nil
	m.Model.BeforeUpdate()
}
