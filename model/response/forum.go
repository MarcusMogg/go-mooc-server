package response

import (
	"server/model/entity"
	"server/model/request"
)

// TopicDetailResp 返回给前端的主题详情
type TopicDetailResp struct {
	Title               string
	Num                 uint // 当前主题的帖子数量
	request.CourseIDReq      // 所属课程
	Posts               []entity.Post
}

// TopicList 主题列表
type TopicList struct {
	Num    uint //总数
	Topics []entity.Topic
}
