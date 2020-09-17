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

// TopicInsResp 返回给前端的主题简介
type TopicInsResp struct {
	Title               string
	Num                 uint // 主题的帖子数量
	request.CourseIDReq      // 所属课程
	UID                 uint // 发布人
}

// TopicList 主题列表
type TopicList struct {
	Num    uint //总数
	Topics []TopicInsResp
}
