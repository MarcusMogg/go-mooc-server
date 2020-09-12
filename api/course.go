package api

import (
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"
	"strconv"

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
		if err = service.InsertCourse(&courseData, user.ID); err == nil {
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

// GetCourseList 获取所有课程
func GetCourseList(c *gin.Context) {
	pagenum, err1 := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	pagesize, err2 := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	keyword := c.Query("key")
	if err1 != nil || err2 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	courses := service.GetCourses(pagenum, pagesize, keyword)
	response.OkWithData(courses, c)
}

// DeleteCourse 删除课程
func DeleteCourse(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var id request.GetByID
	if err := c.BindJSON(&id); err == nil {
		service.DropCourse(id.ID, user.ID)
	} else {
		response.FailValidate(c)
	}
}

// ReadVideoList 读取课程下的视频列表
func ReadVideoList(c *gin.Context) {
	var id request.GetByID
	if err := c.BindJSON(&id); err == nil {
		videos := service.GetVideosByCourseID(id.ID)
		response.OkWithData(videos, c)
	} else {
		response.FailValidate(c)
	}
}

// ReadVideo 读取视频信息
func ReadVideo(c *gin.Context) {
	var id request.GetByID
	if err := c.BindJSON(&id); err == nil {
		video := service.GetVideoByVideoID(id.ID)
		response.OkWithData(video, c)
	} else {
		response.FailValidate(c)
	}
}

//
func ModifyVideoList() {

}

// AddStudents 批量添加学生
func AddStudents(c *gin.Context) {
	var as request.AddStudentsReq
	if err := c.BindJSON(&as); err == nil {
		var errs []string
		for _, i := range as.UserNames {
			if err = service.InsertStudent(as.ID, i); err != nil {
				errs = append(errs, i)
			}
		}
		if len(errs) != 0 {
			response.FailDetailed(response.ERROR, errs, "用户名错误", c)
		} else {
			response.Ok(c)
		}
	} else {
		response.FailValidate(c)
	}
}

// GetStudents 获取用户列表
func GetStudents(c *gin.Context) {
	var id request.GetByID
	if err := c.BindJSON(&id); err == nil {
		var user []entity.MUser = service.GetStudents(id.ID)
		response.OkWithData(user, c)
	} else {
		response.FailValidate(c)
	}
}

// AddWacthTime 增加学生学习时长
func AddWacthTime(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var cs entity.CourseStudents
	if err := c.BindJSON(&cs); err == nil {
		cs.StudentID = user.ID
		service.AddWatchTime(&cs)
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

// GetWatchTimeList 获取所有的观看视频市场
func GetWatchTimeList(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	res := service.GetWatchTimes(user.ID)
	response.OkWithData(res, c)
}
