package service

import (
	"server/global"
	"server/model/entity"
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
func GetIWatchWho(id uint) []int {
	var res []int
	global.GDB.Model(&entity.WatchList{}).Select("to_id").Where("from_id = ?", id).Scan(&res)
	return res
}

// GetWhoWatchI 查询关注我的用户
func GetWhoWatchI(id uint) []int {
	var res []int
	global.GDB.Model(&entity.WatchList{}).Select("from_id").Where("to_id = ?", id).Scan(&res)
	return res
}
