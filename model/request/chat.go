package request

import (
	"time"
)

// CMRType 请求类型
type CMRType int

const (
	// SENDMSG 发送消息
	SENDMSG CMRType = iota + 1
	// ACKMSG 确认收到消息
	ACKMSG
)

// ChatMsgReq 单独聊天时发送的请求
type ChatMsgReq struct {
	CMRType  CMRType   `json:"type"`
	ToID     uint      `json:"id"`
	SendTime time.Time `json:"sendtime"`
	Msg      string    `json:"msg"`
}
