package api

import (
	"fmt"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Apply 申请成为教师
func Apply(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	if user.Role != entity.Student {
		response.FailWithMessage("只有学生可以申请", c)
		return
	}
	a := &entity.ApplyTeacher{
		UserID: user.ID,
		State:  0,
	}
	if err := service.InsertApply(a); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

//GetApply 获取所有教师申请审核
func GetApply(c *gin.Context) {
	pagenum, err1 := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	pagesize, err2 := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	fmt.Println(c.Query("pagesize"))
	if err1 != nil || err2 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	applys, tot, err := service.PaginateApply(pagenum, pagesize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{
		"total": tot,
		"data":  applys,
	}, c)
}

//AgreeApply 审核申请
func AgreeApply(c *gin.Context) {
	var a request.ApplyAgreeReq
	if err := c.BindJSON(&a); err == nil {
		err = service.ChangeApplyState(&a)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.Ok(c)
		}
	} else {
		response.FailValidate(c)
	}
}
