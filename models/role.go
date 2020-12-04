package models

// Role 角色实体
type Role struct {
	//ID        int64     `gorm:"column:id;primary_key;size:36;"`
	Name      *string     `gorm:"column:name;size:100;index;default:'';not null;"` // 角色名称
	Sequence  *int        `gorm:"column:sequence;index;default:0;not null;"`       // 排序值
	Memo      *string    `gorm:"column:varchar;size:200;"`                          // 备注
	Status    *int        `gorm:"column:status;index;default:0;not null;"`         // 状态(1:启用 2:禁用)
	Model
}
