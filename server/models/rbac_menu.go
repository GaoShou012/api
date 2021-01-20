package models

import "api/global"

type RbacMenu struct {
	Model
	GroupId *uint64 `json:",omitempty"`
	Sort    *uint64 `json:",omitempty"`
	Name    *string `json:",omitempty"`
	Code    *string `json:",omitempty"`
	Icon    *string `json:",omitempty"`
	Desc    *string `json:",omitempty"`
}

func (m *RbacMenu) GetTableName() string {
	return "rbac_menu"
}

func (m *RbacMenu) Insert() error {
	res := global.DBMaster.Table(m.GetTableName()).Create(m)
	return res.Error
}

func (m *RbacMenu) SelectByCode(fields string, code string) error {
	res := global.DBSlave.Table(m.GetTableName()).Select(fields).Where("`code`=?", code).First(m)
	return res.Error
}
