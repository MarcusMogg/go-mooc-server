package api

import (
	"fmt"
	"server/middleware"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var r request.RegisterData
	if err := c.BindJSON(&r); err == nil {
		user := &entity.MUser{UserName: r.UserName, Email: r.Email, Password: r.Password, Role: entity.Student}

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
			tokenNext(c, user)
		} else {
			response.FailWithMessage("账号或者密码错误", c)
		}

	} else {
		response.FailValidate(c)
	}
}

func tokenNext(c *gin.Context, u *entity.MUser) {
	j := middleware.NewJWT()
	claim := middleware.JWTClaim{
		UserID:   u.ID,
		UserName: u.UserName,
		Role:     u.Role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 60*60*24*7,
			Issuer:    "715worker",
		},
	}
	token, err := j.CreateToken(claim)
	if err != nil {
		response.FailWithMessage("token创建失败", c)
		return
	}
	response.OkWithData(token, c)
}
