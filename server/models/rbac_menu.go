package models

type RbacMenu struct {
	Model
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

