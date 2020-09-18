package service

import (
	"server/global"
	"server/model/entity"

	"gorm.io/gorm"
)

// InsertLive 添加直播
func InsertLive(live *entity.Live) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(live).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetLiveByID 通过课程id获取课程信息
func GetLiveByID(id uint) *entity.Live {
	var l entity.Live
	global.GDB.Where("id = ?", id).First(&l)
	return &l
}

// GetLiveByCourseID 通过教师id获取课程列表
func GetLiveByCourseID(id uint) []entity.Live {
	var l []entity.Live
	global.GDB.Where("course_id = ?", id).Find(&l)
	return l
}

// InsertLiveSign 插入签到信息
func InsertLiveSign(liveSign *entity.LiveSign) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(liveSign).Error
	})
}

// ReadLiveSign 读取全部签到信息
func ReadLiveSign(liveID uint) []entity.LiveSign {
	var s []entity.LiveSign
	global.GDB.Where("id = ?", liveID).Find(&s)
	return s
}
