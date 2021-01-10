package models

//系统设置
type SettingSys struct {
	Model
}

func (m *SettingSys) GetTableName() string {
	return "setting_sys"
}
