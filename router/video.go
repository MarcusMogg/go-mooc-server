package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

//InitVideoRouter 初始化video路由组
func InitVideoRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("video")
	{
		UserRouter.POST("upload", api.Upload)
	}
}
