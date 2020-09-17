package api

import (
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"

	"github.com/gin-gonic/gin"
)

// CreateTopic 创建主题
func CreateTopic(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var t request.TopicReq
	if err := c.BindJSON(&t); err == nil {
		topic := &entity.Topic{
			CID:       t.CourseIDReq.ID,
			UserID:    user.ID,
			Title:     t.Title,
			Top:       false,
			Important: false,
		}
		post := &entity.Post{
			UserID:   user.ID,
			Msg:      t.Msg,
			ReplyFor: 0,
		}
		service.InsertTopic(topic, post)
		response.OkWithData(topic.ID, c)
	} else {
		response.FailValidate(c)
	}
}

// CreatePost 创建帖子
func CreatePost(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var t request.PostReq
	if err := c.BindJSON(&t); err == nil {
		post := &entity.Post{
			UserID:   user.ID,
			Msg:      t.Msg,
			ReplyFor: t.Reply,
			TopicID:  t.TopicID,
		}
		service.InsertPost(post)
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// GetTopicList 获取当前课程相关的主题
func GetTopicList(c *gin.Context) {
	var tr request.GetTopicsReq
	if err := c.BindJSON(&tr); err == nil {
		if tr.Pagenum == 0 {
			tr.Pagenum = 1
		}
		if tr.Pagesize == 0 {
			tr.Pagesize = 20
		}

		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// GetTopicDetail 获取主题详情
func GetTopicDetail(c *gin.Context) {

}
