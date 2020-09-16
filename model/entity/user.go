package entity

import "gorm.io/gorm"

// MUser 数据库用户字段
type MUser struct {
	gorm.Model
	UserName string `gorm:"not null;unique"`
	Email    string
	NickName string
	Password string `gorm:"not null" json:"-"`
	Avatar   string // 用户头像地址
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

// WatchList from关注to
type WatchList struct {
	FromID uint `gorm:"primaryKey;autoIncrement:false"`
	ToID   uint `gorm:"primaryKey;autoIncrement:false"`
}
