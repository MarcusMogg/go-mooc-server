package entity

import (
	"gorm.io/gorm"
)

// Topic 帖子
type Topic struct {
	gorm.Model
	CreatedAtStr string `gorm:"-" json:"time"`
	UserID       uint
	CID          uint `gorm:"not null"` //关联的课程ID
	Num          uint // 主题的帖子数量
	Title        string
	Top          bool // 置顶
	Important    bool // 加精
}

// Post 回复
type Post struct {
	gorm.Model
	CreatedAtStr string `gorm:"-" json:"time"`
	TopicID      uint   `json:"-"`
	UserID       uint
	ReplyFor     uint
	Msg          string
	Num          uint // 点赞数量
}

// TopicAuth 权限
type TopicAuth uint

const (
	// POST 发帖
	POST TopicAuth = 1 << iota
	// TOP 置顶
	TOP
	// IMPORTANT 加精
	IMPORTANT
	// DELETE 删帖
	DELETE
	// APPROVE 审核学生
	APPROVE
)

// CheckTopicAuth 检查是否拥有相应权限
func CheckTopicAuth(auth, target TopicAuth) bool {
	return (auth | target) > 0
}

// SetTopicAuth 设置相应权限
func SetTopicAuth(auth, target TopicAuth) TopicAuth {
	return auth | target
}

// UserLike 用户点赞记录
type UserLike struct {
	UID uint `gorm:"primaryKey;autoIncrement:false"`
	PID uint `gorm:"primaryKey;autoIncrement:false"`
}
