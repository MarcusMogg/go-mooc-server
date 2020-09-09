package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

//InitUserRouter 初始化user路由组
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", api.Register)
		UserRouter.POST("login", api.Login)
	}
}
