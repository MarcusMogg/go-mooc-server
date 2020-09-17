package entity

import (
	"gorm.io/gorm"
)

// Course 课程信息
type Course struct {
	gorm.Model
	TeacherID   uint
	Name        string
	Instruction string
}

// CourseStudents 加入课程的学生
type CourseStudents struct {
	StudentID uint      `gorm:"primaryKey;autoIncrement:false" json:"-"`
	CourseID  uint      `gorm:"primaryKey;autoIncrement:false" json:"cid"`
	WatchTime uint64    `json:"watchtime"`
	Status    uint      `json:"-"` // 0 申请中 1 已加入
	Auth      TopicAuth `json:"-"`
}
