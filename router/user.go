package router

import (
	"server/api"
	"server/middleware"
	"server/model/entity"

	"github.com/gin-gonic/gin"
)

//InitUserRouter 初始化user路由组
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", api.Register)
		UserRouter.POST("login", api.Login)
		UserRouter.GET("apply", middleware.JWTAuth(), middleware.RoleAuth(entity.Admin), api.GetApply)
		UserRouter.POST("apply", middleware.JWTAuth(), api.Apply)
		UserRouter.POST("agree", middleware.JWTAuth(), middleware.RoleAuth(entity.Admin), api.AgreeApply)
	}
}
