package initialize

import (
	"server/router"

	"github.com/gin-gonic/gin"
)

// Router 初始化路由列表
func Router() *gin.Engine {
	var Router = gin.Default()

	APIGroup := Router.Group("")
	router.InitUserRouter(APIGroup)

	return Router
}
