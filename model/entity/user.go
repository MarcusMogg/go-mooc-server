package entity

import "gorm.io/gorm"

// MUser 数据库用户字段
type MUser struct {
	gorm.Model
	UserName string `gorm:"not null;unique"`
	Email    string
	NickName string
	Password string `gorm:"not null" json:"-"`
	Role     Role
}

// Role 用户身份
type Role int

const (
	// Student 学生
	Student Role = iota
	// Teacher 老师
	Teacher
	// Admin 管理员
	Admin
)

// FriendRequest 好友申请数据库
type FriendRequest struct {
	FromID uint `gorm:"primaryKey;autoIncrement:false"`
	ToID   uint `gorm:"primaryKey;autoIncrement:false"`
	Status bool // 同意与否
}

// UserFriend 好友信息数据库
type UserFriend struct {
	UserID   uint `gorm:"primaryKey;autoIncrement:false"`
	FriendID uint `gorm:"primaryKey;autoIncrement:false"`
}
