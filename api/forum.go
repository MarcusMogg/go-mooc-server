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
			Num:       0,
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
			Num:      0,
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
			Num:      0,
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
		res := service.GetTopicsByCourseID(tr.Pagenum, tr.Pagesize, tr.CourseIDReq.ID)

		response.OkWithData(res, c)
	} else {
		response.FailValidate(c)
	}
}

// GetTopicDetail 获取主题详情
func GetTopicDetail(c *gin.Context) {
	var tr request.GetTopicDetailReq
	if err := c.BindJSON(&tr); err == nil {
		if tr.Pagenum == 0 {
			tr.Pagenum = 1
		}
		if tr.Pagesize == 0 {
			tr.Pagesize = 20
		}
		res := service.GetTopicDetail(tr.Pagenum, tr.Pagesize, tr.GetByID.ID)

		response.OkWithData(res, c)
	} else {
		response.FailValidate(c)
	}
}

// DeleteTopic 删除主题
func DeleteTopic(c *gin.Context) {
	var id request.CommonTopicReq
	if err := c.BindJSON(&id); err == nil {
		service.DropTopic(id.GetByID.ID)
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// DeletePost 删除帖子
func DeletePost(c *gin.Context) {
	var id request.CommonTopicReq
	if err := c.BindJSON(&id); err == nil {
		service.DropPost(id.GetByID.ID)
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// TopTopic 置顶帖子
func TopTopic(c *gin.Context) {
	var req request.CommonTopicReq
	if err := c.BindJSON(&req); err == nil {
		service.SetTopicTop(req.GetByID.ID, req.Status)
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// ImportTopic 加精帖子
func ImportTopic(c *gin.Context) {
	var req request.CommonTopicReq
	if err := c.BindJSON(&req); err == nil {
		service.SetTopicImport(req.GetByID.ID, req.Status)
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

//LikeTopic 点赞帖子
func LikeTopic(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var req request.CommonTopicReq
	if err := c.BindJSON(&req); err == nil {
		if req.Status {
			service.Like(user.ID, req.GetByID.ID)
		} else {
			service.UnLike(user.ID, req.GetByID.ID)
		}
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}
