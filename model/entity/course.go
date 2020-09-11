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
