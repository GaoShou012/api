package models

type Admins struct {
	Model
	Username *string `gorm:"type:varchar(20);not null;unique"`
	Password *string `gorm:"type:varchar(150);not null"`
	Nickname *string `gorm:"type:varchar(20)"`
	Status   *int    `json:"status,omitempty" gorm:"column:status;index;default:0;not null;"` // 状态(1:启用 2:停用)
	RoleId   *int64  `json:"role_id,omitempty"`
}

func (a *Admins) SelectByUsername() {
}
