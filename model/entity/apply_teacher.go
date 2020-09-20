package entity

import (
	"gorm.io/gorm"
)

// ApplyTeacher 教师申请列表
type ApplyTeacher struct {
	gorm.Model
	UserID uint `gorm:"not null"` //申请人
	State  int  // 处理状态 0 未处理 1 已通过 2 未通过
}
