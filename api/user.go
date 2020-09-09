package api

import (
	"fmt"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var r request.RegisterData
	if err := c.BindJSON(&r); err == nil {
		user := &entity.MUser{UserName: r.UserName, NickName: r.NickName, Password: r.Password}

		if err = service.Register(user); err == nil {
			response.OkWithMessage("注册成功", c)
		} else {
			response.FailWithMessage(fmt.Sprintf("%v", err), c)
		}

	} else {
		response.FailValidate(c)
	}
}

// Login 用户登录
func Login(c *gin.Context) {
	var r request.LoginData
	if err := c.BindJSON(&r); err == nil {
		user := &entity.MUser{UserName: r.UserName, Password: r.Password}

		if service.Login(user) {
			response.OkDetailed(user, "登录成功", c)
		} else {
			response.FailWithMessage("账号或者密码错误", c)
		}

	} else {
		response.FailValidate(c)
	}
}
