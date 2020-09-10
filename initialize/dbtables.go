package initialize

import (
	"server/global"
	"server/model/entity"
)

// DBTables 迁移 schema
func DBTables() {
	global.GDB.AutoMigrate(&entity.MUser{})
	global
}
