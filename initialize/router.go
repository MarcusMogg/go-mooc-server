package initialize

import (
	"server/middleware"
	"server/router"

	"github.com/gin-gonic/gin"
)

// Router 初始化路由列表
func Router() *gin.Engine {
	var Router = gin.Default()

	Router.Use(middleware.Cors()) // 跨域

	APIGroup := Router.Group("")
	router.InitUserRouter(APIGroup)
	router.InitVideoRouter(APIGroup)

	router.InitCourseRouter(APIGroup)
	router.InitChatRouter(APIGroup)
	router.InitLiveRouter(APIGroup)
	return Router
}
