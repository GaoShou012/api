package models

//对话设置
type SettingChats struct {
	Model
	AssignRule          *int  //分配规则 1按饱和度分配，2依次分配
	IsAssignLastService *bool //是否优先分配上一次对话过的客服
}

func (m *SettingChats) GetTableName() string {
	return "setting_chats"
}
