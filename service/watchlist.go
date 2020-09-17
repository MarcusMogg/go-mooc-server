package service

import (
	"server/global"
	"server/model/entity"
	"server/model/response"
)

// InsertWatchRecord 添加关注记录
func InsertWatchRecord(from, to uint) {
	w := entity.WatchList{
		FromID: from,
		ToID:   to,
	}
	global.GDB.Model(&entity.WatchList{}).Create(&w)
}

// GetIWatchWho 查询我关注的用户
func GetIWatchWho(id uint) []response.IDName {
	var res []response.IDName
	global.GDB.Table("watch_lists").Select("watch_lists.to_id,m_users.nick_name").Joins("JOIN m_users ON m_users.ID = watch_lists.to_id").Where("from_id = ?", id).Scan(&res)
	return res
}

// GetWhoWatchI 查询关注我的用户
func GetWhoWatchI(id uint) []response.IDName {
	var res []response.IDName
	global.GDB.Table("watch_lists").Select("watch_lists.from_id,m_users.nick_name").Joins("JOIN m_users ON m_users.ID = watch_lists.from_id").Where("to_id = ?", id).Scan(&res)
	return res
}
