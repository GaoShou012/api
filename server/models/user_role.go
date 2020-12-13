package models

// UserRole 用户角色关联实体
type UserRole struct {
	ID     int64 `gorm:"column:id;primary_key;"`
	UserID int64 `gorm:"column:user_id;"` // 用户内码
	RoleID int64 `gorm:"column:role_id;"` // 角色内码
}
