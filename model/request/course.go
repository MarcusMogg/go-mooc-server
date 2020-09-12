package request

// CourseReq 创建课程申请
type CourseReq struct {
	//TeacherID   uint   `json:"id" binding:"required"`
	Instruction string `json:"instruction" binding:"required"`
	Name        string `json:"name" binding:"required"`
}

// CourseUReq 课程更新申请
type CourseUReq struct {
	GetByID
	CourseReq
}

// AddStudentsReq 添加学生申请
type AddStudentsReq struct {
	GetByID
	UserNames []string `json:"studentnames" binding:"required"`
}
