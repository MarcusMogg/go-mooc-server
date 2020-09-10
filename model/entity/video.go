package entity

import "gorm.io/gorm"

// Video 数据库用户字段
type Video struct {
	gorm.Model
	VideoName	string
	Course		string
	Uploader	string
	Format		string
	Path 		string
}
