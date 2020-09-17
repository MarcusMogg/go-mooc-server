package response

import (
	"server/model/entity"
	"time"
)

// ChatMsgResp 单独聊天时回复的内容
type ChatMsgResp struct {
	FromID      uint           `json:"from"`
	ToID        uint           `json:"-"`
	SendTime    time.Time      `json:"-" gorm:"column:created_at"`
	SendTimeStr string         `json:"sendtime"`
	Msg         string         `json:"msg"`
	MType       entity.MsgType `json:"msgtype"`
}

// UnreadMsgNumResp 未读消息数量
type UnreadMsgNumResp struct {
	FromID uint `json:"from"`
	Num    uint `json:"num"`
}

// UnreadMsgResp 未读消息
type UnreadMsgResp struct {
	FromID uint          `json:"from"`
	Num    int64         `json:"num"`
	Msg    []ChatMsgResp `json:"msgs"`
}
