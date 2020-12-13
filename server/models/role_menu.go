package models
// RoleMenu 角色菜单实体
type RoleMenu struct {
	ID       int64 `gorm:"column:id;primary_key;"`
	RoleID   string `gorm:"column:role_id;size:36;index;default:'';not null;"`   // 角色ID
	MenuID   string `gorm:"column:menu_id;size:36;index;default:'';not null;"`   // 菜单ID
	ActionID string `gorm:"column:action_id;size:36;index;default:'';not null;"` // 动作ID
}
