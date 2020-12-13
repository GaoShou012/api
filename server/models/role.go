package models

// Role 角色实体
type Role struct {
	ID        *uint64    `json:"id,omitempty" gorm:"primary_key"`
	Name      *string     `gorm:"column:name;size:100;index;default:'';not null;"` // 角色名称
	Sequence  *int        `gorm:"column:sequence;index;default:0;not null;"`       // 排序值
	Memo      *string    `gorm:"column:memo;size:200;"`                          // 备注
	Status    *int        `gorm:"column:status;index;default:1;not null;"`         // 状态(1:启用 2:禁用)
}