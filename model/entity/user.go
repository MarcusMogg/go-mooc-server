package entity

import "gorm.io/gorm"

type MUser struct {
	gorm.Model
	UserName string `gorm:"not null;unique"`
	NickName string
	Password string
}
