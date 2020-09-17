package request

// PostReq 发帖请求
type PostReq struct {
	CourseIDReq
	TopicID uint `json:"topic"`
	Reply   uint
	Msg     string
}

// TopicReq 发帖请求
type TopicReq struct {
	CourseIDReq
	Title string
	Msg   string
}

// CommonTopicReq 加精\置顶\删除
type CommonTopicReq struct {
	CourseIDReq
	GetByID
	Status bool
}

// GetTopicsReq 获取主题列表
type GetTopicsReq struct {
	Pagenum  int
	Pagesize int
	CourseIDReq
}

// GetTopicDetailReq 获取主题详情
type GetTopicDetailReq struct {
	Pagenum  int
	Pagesize int
	GetByID
	CourseIDReq
}
