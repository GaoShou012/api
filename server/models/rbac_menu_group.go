package models

import "api/global"

type RbacMenuGroup struct {
	Model
	Sort *uint64
	Name *string
	Code *string
	Icon *string
	Desc *string
}

func (m *RbacMenuGroup) GetTableName() string {
	return "rbac_menu_group"
}

func (m *RbacMenuGroup) SelectByCode(fields string, code string) error {
	res := global.DBSlave.Table(m.GetTableName()).Select(fields).Where("code=?", code).First(m)
	return res.Error
}

func (m *RbacMenuGroup) IsExistsByCode(code string) (bool, error) {
	count := 0
	res := global.DBSlave.Table(m.GetTableName()).Where("code=?", code).Count(&count)
	if res.Error != nil {
		return false, res.Error
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (m *RbacMenuGroup) Insert() error {
	res := global.DBMaster.Table(m.GetTableName()).Create(m)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
