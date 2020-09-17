package initialize

import (
	"server/global"
	"server/model/entity"
	"server/service"
)

// DBTables 迁移 schema
func DBTables() {
	global.GDB.AutoMigrate(&entity.MUser{})
	global.GDB.AutoMigrate(&entity.Video{})
	u := entity.MUser{
		UserName: "admin",
		Password: "123456",
		Role:     entity.Admin,
	}
	service.Register(&u)
	global.GDB.AutoMigrate(&entity.ApplyTeacher{})
	global.GDB.AutoMigrate(&entity.Course{})
	global.GDB.AutoMigrate(&entity.CourseStudents{})
	global.GDB.AutoMigrate(&entity.ChatMessage{})
	global.GDB.AutoMigrate(&entity.FriendRequest{})
	global.GDB.AutoMigrate(&entity.UserFriend{})
	global.GDB.AutoMigrate(&entity.Live{})
}
