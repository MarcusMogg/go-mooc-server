package router

import (
	"server/api"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

//InitChatRouter 初始化chat路由组
func InitChatRouter(Router *gin.RouterGroup) {
	ChatRouter := Router.Group("chat")
	{
		ChatRouter.GET("ws", middleware.JWTAuth(), api.AloneWS)
		ChatRouter.GET("unread", middleware.JWTAuth(), api.GetUnreadMsg)
		ChatRouter.GET("unreadnum", middleware.JWTAuth(), api.GetUnreadMsgNum)
	}
}
