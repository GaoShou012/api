package models

// MenuActionResource 菜单动作关联资源实体
type MenuActionResource struct {
	ID       int64  `gorm:"column:id;"`
	ActionID int64  `gorm:"column:action_id;"`                           // 菜单动作ID
	Method   string `gorm:"column:method;size:100;default:'';not null;"` // 资源请求方式(支持正则)
	Path     string `gorm:"column:path;size:100;default:'';not null;"`   // 资源请求路径（支持/:id匹配）
}
