package models

type Admins struct {
	Model
	Username *string `gorm:"type:varchar(20);not null;unique"`
	Password *string `gorm:"type:varchar(255);not null"`
	Nickname *string `gorm:"type:varchar(255);not null"`
}
