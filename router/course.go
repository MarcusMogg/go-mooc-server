package router

import (
	"server/api"
	"server/middleware"
	"server/model/entity"

	"github.com/gin-gonic/gin"
)

//InitCourseRouter 初始化course路由组
func InitCourseRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("course")
	{
		UserRouter.POST("create", middleware.JWTAuth(), middleware.RoleAuth(entity.Teacher), api.CreateCourse)
		UserRouter.POST("update", middleware.JWTAuth(), middleware.RoleAuth(entity.Teacher), api.UpdateCourse)
		UserRouter.POST("get", api.ReadCourse)
		UserRouter.POST("getlist", middleware.JWTAuth(), api.ReadCourseList)
		UserRouter.POST("getvideolist", api.ReadVideoList)
		UserRouter.POST("getvideo", api.ReadVideo)
	}
}
