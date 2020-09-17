package request

// LiveReqType 请求类型
type LiveReqType int

const (
	// MSG 发送消息
	MSG LiveReqType = iota + 1
	CTR
)

// LiveReq 直播信息
type LiveReq struct {
	Name         string `form:"name" json:"name" binding:"required"`
	Introduction string `form:"introduction" json:"introduction" binding:"required"`
	CourseID     uint   `form:"courseId" json:"courseId" binding:"required"`
	CourseName   string `form:"courseId" json:"courseName" binding:"required"`
	StartTime    string `form:"startTime" json:"startTime" binding:"required"`
	EndTime      string `form:"endTime" json:"endTime" binding:"required"`
}

/* UserSigReq 直播密钥请求
type UserSigReq struct {
	sdkappid int    `form:"sdkappid" json:"sdkappid" binding:"required"`
	username string `form:"username" json:"username" binding:"required"`
}*/

// LiveMsgReq 聊天室消息
type LiveMsgReq struct {
	LiveReqType LiveReqType `json:"type"`
	ChatData    ChatData
	ControlData ControlData
}

// ControlData 控制主体
type ControlData struct {
}

// IDData ID信息
type IDData struct {
}

// ChatData 消息主体
type ChatData struct {
	Text TextBody
	Mine bool   `json:"mine"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

// TextBody 文本
type TextBody struct {
	Text string `json:"text"`
}

// EnterRoom 进去直播间需要的信息
type EnterRoom struct {
	IsTeacher bool   `json:"isteacher"`
	UName     string `json:"uname"`
	Icon      string `json:"icon"`
	LiveID    uint   `json:"liveid"`
	UID       uint   `json:"uid"`
}
