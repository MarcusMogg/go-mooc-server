package router

import (
	"server/api"
	"server/middleware"
	"server/model/entity"

	"github.com/gin-gonic/gin"
)

//InitForumRouter 初始化forum路由组
func InitForumRouter(Router *gin.RouterGroup) {
	forumRouter := Router.Group("forum")
	{
		forumRouter.POST("create", middleware.JWTAuth(), middleware.TopicAuth(entity.POST), api.CreateTopic)
		forumRouter.POST("reply", middleware.JWTAuth(), middleware.TopicAuth(entity.POST), api.CreatePost)
		forumRouter.POST("list", middleware.JWTAuth(), middleware.TopicAuth(entity.POST), api.GetTopicList)
		forumRouter.POST("detail", middleware.JWTAuth(), middleware.TopicAuth(entity.POST), api.GetTopicDetail)
		forumRouter.POST("like", middleware.JWTAuth(), middleware.TopicAuth(entity.POST), api.LikeTopic)
		forumRouter.POST("top", middleware.JWTAuth(), middleware.TopicAuth(entity.TOP), api.TopTopic)
		forumRouter.POST("import", middleware.JWTAuth(), middleware.TopicAuth(entity.IMPORTANT), api.ImportTopic)
		forumRouter.POST("deletetopic", middleware.JWTAuth(), middleware.TopicAuth(entity.DELETE), api.DeleteTopic)
		forumRouter.POST("deletepost", middleware.JWTAuth(), middleware.TopicAuth(entity.DELETE), api.DeletePost)
	}
}
