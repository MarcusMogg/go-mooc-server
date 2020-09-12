package api

import (
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"

	"github.com/gin-gonic/gin"
)

// CreateCourse 创建课程
func CreateCourse(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var course request.CourseReq
	if err := c.BindJSON(&course); err == nil {
		courseData := entity.Course{
			TeacherID:   user.ID,
			Instruction: course.Instruction,
			Name:        course.Name,
		}
		if err = service.InsertCourse(&courseData); err == nil {
			response.Ok(c)
		} else {
			response.FailWithMessage("课程创建失败", c)
		}
	} else {
		response.FailValidate(c)
	}

}

// UpdateCourse 修改课程信息
func UpdateCourse(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var cs request.CourseUReq
	if err := c.BindJSON(&cs); err == nil {
		course := service.GetCourseByID(cs.ID)
		course.Instruction = cs.Instruction
		course.Name = cs.Name
		if err = service.UpdateCourse(course, user); err == nil {
			response.Ok(c)
		} else {
			response.FailWithMessage(err.Error(), c)
		}
	} else {
		response.FailValidate(c)
	}

}

// ReadCourse 读取课程信息
func ReadCourse(c *gin.Context) {
	var id request.GetByID
	if err := c.BindJSON(&id); err == nil {
		course := service.GetCourseByID(id.ID)
		response.OkWithData(course, c)
	} else {
		response.FailValidate(c)
	}
}

// ReadCourseList 读取教师创建的课程列表
func ReadCourseList(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	courses := service.GetCoursesByTeacID(user.ID)
	response.OkWithData(courses, c)
}

// DeleteCourse 删除课程
func DeleteCourse(c *gin.Context) {

}

// ReadVideoList 读取课程下的视频列表
func ReadVideoList(c *gin.Context) {

}

//
func ReadVideoPath() {

}
