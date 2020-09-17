package request

import (
	"server/model/entity"
)

// CourseIDReq 课程ID
type CourseIDReq struct {
	ID uint `json:"cid" form:"cid" binding:"required"`
}

// CourseReq 创建课程申请
type CourseReq struct {
	//TeacherID   uint   `json:"id" binding:"required"`
	Instruction string `json:"instruction" binding:"required"`
	Name        string `json:"name" binding:"required"`
}

// CourseUReq 课程更新申请
type CourseUReq struct {
	CourseIDReq
	CourseReq
}

// AddStudentsReq 添加学生申请
type AddStudentsReq struct {
	CourseIDReq
	UserNames []string `json:"studentnames" binding:"required"`
}

// ApproveStudentApplyReq 通过学生申请
type ApproveStudentApplyReq struct {
	CourseIDReq
	GetByID
	Status bool // true 同意, false 拒绝
}

// StudentAuthReq 学生权限
type StudentAuthReq struct {
	CourseIDReq
	GetByID
	Auth entity.TopicAuth
}
