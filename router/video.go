package router

import (
	"server/api"
	"server/middleware"
	"server/model/entity"

	"github.com/gin-gonic/gin"
)

//InitVideoRouter 初始化video路由组
func InitVideoRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("video")
	{
		UserRouter.POST("upload", middleware.JWTAuth(), middleware.RoleAuth(entity.Teacher), api.Upload)
	}
}
