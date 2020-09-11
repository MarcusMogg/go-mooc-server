package entity

import "gorm.io/gorm"

// Video 视频
type Video struct {
	gorm.Model
	VideoName	string
	Name		string
	Format		string
	Seq			int
	Path 		string
}

// CourseVideo 关联课程视频
type CourseVideo struct {
	gorm.Model
	CourseID	uint
	VideoID		uint
}

// CourseVideoResult 关联查询结果
type CourseVideoResult struct {
	CourseID uint
	Seq 	 int
}