package response

// LiveResType 返回类型
type LiveResType int

const (
	// MSG 发送消息
	MSG LiveResType = iota + 1
	ETR
)

// EnterRoom 进入直播间返回消息
type EnterRoom struct {
	Type      LiveResType `json:"type"`
	UID       uint        `json:"uid"`
	UName     string      `json:"uname"`
	Icon      string      `json:"icon"`
	IsTeacher bool        `json:"isteacher"`
}
