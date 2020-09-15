package entity

import "gorm.io/gorm"

// Video 视频
type Video struct {
	gorm.Model

	// 查询视频重复
	CourseID uint
	Seq      uint

	// 获取静态路径 Path+VideoName+Format
	Path      string
	VideoName string
	Format    string

	// 视频展示
	Name         string
	Introduction string
}
