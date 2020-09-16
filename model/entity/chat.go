package entity

import (
	"gorm.io/gorm"
)

const (
	// READ 消息已读
	READ = true
	// UNREAD 消息未读
	UNREAD = false
)

// MsgType 消息类型
type MsgType uint

const (
	// MAlone 单人聊天
	MAlone MsgType = iota + 1
	// MBroadcast 广播
	MBroadcast
	// MReply 回复
	MReply
	// MLike 点赞
	MLike
	// MFriendReq 好友申请
	MFriendReq
)

// ChatMessage 消息数据库
type ChatMessage struct {
	gorm.Model
	FromID uint // 发送人
	ToID   uint // 接收人
	Msg    string
	Status bool // 是否已读
	MType  MsgType
}
