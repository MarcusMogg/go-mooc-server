package entity

import "gorm.io/gorm"

// Video 视频
type Video struct {
	gorm.Model
	VideoName	string
	Name		string
	Format		string
	Order		int
	Path 		string
}

// CourseVideo 关联课程视频
type CourseVideo struct {
	gorm.Model
	CourseID	int
	VideoID		int
}