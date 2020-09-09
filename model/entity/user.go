package entity

import "gorm.io/gorm"

// MUser 数据库用户字段
type MUser struct {
	gorm.Model
	UserName string `gorm:"not null;unique"`
	Email    string
	NickName string
	Password string
}
