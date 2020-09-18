package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

//InitLiveRouter 初始化live路由组
func InitLiveRouter(Router *gin.RouterGroup) {
	LiveRouter := Router.Group("live")
	{
		LiveRouter.GET("ws", api.LiveWS)
		LiveRouter.POST("key", api.GetUserSig)
	}
}
