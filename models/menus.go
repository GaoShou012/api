package models

import "time"

// Menu 菜单实体
type Menus struct {
	Id        *uint64
	Name      *string // 菜单名称
	Sort      *int    // 排序值
	Icon      *string // 菜单图标
	Path      *string // 访问路由
	Creator   *string // 创建人
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
