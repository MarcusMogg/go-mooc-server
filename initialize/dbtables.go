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
	global.GDB.AutoMigrate(&entity.CourseVideo{})
	u := entity.MUser{
		UserName: "admin",
		Password: "123456",
		Role:     entity.Admin,
	}
	service.Register(&u)
	global.GDB.AutoMigrate(&entity.ApplyTeacher{})
	global.GDB.AutoMigrate(&entity.Course{})
}
