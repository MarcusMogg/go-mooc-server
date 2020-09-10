package entity

import "gorm.io/gorm"

// Video 数据库用户字段
type Video struct {
	gorm.Model
	videoName	string
	course		string 
	status		string
	path 		string
}
