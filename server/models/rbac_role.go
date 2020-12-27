package models

type RbacRole struct {
	Model
	Name     *string
	Desc     *string
	Icon     *string
}

func (m *RbacRole) GetTableName() string {
	return "rbac_role"
}

