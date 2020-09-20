package api

import (
	"fmt"
	"server/global"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"

	"github.com/gin-gonic/gin"
)

// AddStudents 批量添加学生
func AddStudents(c *gin.Context) {
	var as request.AddStudentsReq
	if err := c.BindJSON(&as); err == nil {
		var errs []string
		for _, i := range as.UserNames {
			if id, err := service.InsertStudent(as.ID, i, 1); err != nil {
				errs = append(errs, i)
			} else {
				sendMessage(0, id, fmt.Sprintf("您已加入课程%d", as.ID), entity.MBroadcast)
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
	var id request.CourseIDReq
	if err := c.BindJSON(&id); err == nil {
		var user []entity.MUser = service.GetStudents(id.ID, 1)
		response.OkWithData(user, c)
	} else {
		response.FailValidate(c)
	}
}

// GetApplyStudents 获取申请中的用户列表
func GetApplyStudents(c *gin.Context) {
	var id request.CourseIDReq
	if err := c.BindJSON(&id); err == nil {
		var user []entity.MUser = service.GetStudents(id.ID, 0)
		response.OkWithData(user, c)
	} else {
		response.FailValidate(c)
	}
}

//ApplyCourse 学生申请课程
func ApplyCourse(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var id request.CourseIDReq
	if err := c.BindJSON(&id); err == nil {
		if uid, err := service.InsertStudent(id.ID, user.UserName, 0); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			sendMessage(0, uid, fmt.Sprintf("您已申请加入课程%d", id.ID), entity.MBroadcast)
			response.Ok(c)
		}
	} else {
		response.FailValidate(c)
	}
}

// ApproveCourseApply 教师通过/拒绝学生申请
func ApproveCourseApply(c *gin.Context) {
	var a request.ApproveStudentApplyReq
	if err := c.BindJSON(&a); err == nil {
		if a.Status { // 同意更新学生状态
			service.UpdateStudentStatus(a.GetByID.ID, a.CourseIDReq.ID, 1)
			sendMessage(0, a.GetByID.ID, "已通过课程申请", entity.MBroadcast)
		} else { // 失败删除
			service.DeleteStudent(a.GetByID.ID, a.CourseIDReq.ID)
			sendMessage(0, a.GetByID.ID, "未通过课程申请", entity.MBroadcast)
		}
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// DeleteStudent 删除学生
func DeleteStudent(c *gin.Context) {
	var a request.ApproveStudentApplyReq
	if err := c.BindJSON(&a); err == nil {
		err := service.DeleteStudent(a.GetByID.ID, a.CourseIDReq.ID)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			sendMessage(0, a.GetByID.ID, fmt.Sprintf("您已被移出课程%d", a.CourseIDReq.ID), entity.MBroadcast)
			response.Ok(c)
		}
	} else {
		response.FailValidate(c)
	}
}

// UpdateStudentAuth 修改学生权限
func UpdateStudentAuth(c *gin.Context) {
	var id request.StudentAuthReq
	if err := c.BindJSON(&id); err == nil {
		service.SetStudentAuth(id.GetByID.ID, id.CourseIDReq.ID, id.Auth)
		sendMessage(0, id.GetByID.ID, "您的权限已被修改，请重新登录", entity.MBroadcast)
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// GetStudentsAuth 获取用户权限
func GetStudentsAuth(c *gin.Context) {
	var id request.StudentAuthReq
	if err := c.BindJSON(&id); err == nil {
		res := service.GetStudentAuth(id.GetByID.ID, id.CourseIDReq.ID)
		response.OkWithData(res, c)
	} else {
		response.FailValidate(c)
	}
}

// InCourse 判断学生是否在课程里
func InCourse(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var id request.CourseIDReq
	if err := c.BindJSON(&id); err == nil {
		if err := service.CourseExist(id.ID); err == nil {
			response.OkWithData(service.StudentInCourse(id.ID, user.ID, global.GDB), c)
			return
		}
	} else {
		response.FailValidate(c)
	}
	response.Fail(c)
}
