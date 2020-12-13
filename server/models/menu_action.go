package models

type MenuAction struct {
	ID     int64  `gorm:"column:id;primary_key;"`
	MenuID string `gorm:"column:menu_id;size:36;index;default:'';not null;"` // 菜单ID
	Code   string `gorm:"column:code;size:100;default:'';not null;"`         // 动作编号
	Name   string `gorm:"column:name;size:100;default:'';not null;"`         // 动作名称
}
