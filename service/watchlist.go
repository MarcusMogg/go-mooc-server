package service

import (
	"errors"
	"server/global"
	"server/model/entity"
	"server/model/response"

	"gorm.io/gorm"
)

// InsertWatchRecord 添加关注记录
func InsertWatchRecord(from, to uint) error {
	w := entity.WatchList{
		FromID: from,
		ToID:   to,
	}
	return global.GDB.Model(&entity.WatchList{}).Create(&w).Error
}

// GetIWatchWho 查询我关注的用户
func GetIWatchWho(id uint) []response.IDName {
	var res []response.IDName
	global.GDB.Table("watch_lists").Select("watch_lists.to_id as id,m_users.nick_name").Joins("JOIN m_users ON m_users.ID = watch_lists.to_id").Where("from_id = ?", id).Scan(&res)
	return res
}

// GetWhoWatchI 查询关注我的用户
func GetWhoWatchI(id uint) []response.IDName {
	var res []response.IDName
	global.GDB.Table("watch_lists").Select("watch_lists.from_id as id,m_users.nick_name").Joins("JOIN m_users ON m_users.ID = watch_lists.from_id").Where("to_id = ?", id).Scan(&res)
	return res
}

// IsWatchWho 查询我关注的用户
func IsWatchWho(from uint, to uint) bool {
	result := global.GDB.Where("from_id = ? AND to_id = ?", from, to).First(&entity.WatchList{})
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// DropWatchWho 删除关注的用户
func DropWatchWho(from uint, to uint) {
	global.GDB.Where("from_id = ? AND to_id = ?", from, to).Delete(&entity.WatchList{})
}
