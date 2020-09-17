package service

import (
	"server/global"
	"server/model/entity"
	"server/model/response"

	"gorm.io/gorm"
)

// InsertTopic 创建主题
func InsertTopic(t *entity.Topic, p *entity.Post) {
	global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(t).Error; err != nil {
			return err
		}
		p.TopicID = t.ID
		return tx.Create(p).Error
	})
}

// InsertPost 创建帖子
func InsertPost(p *entity.Post) {
	global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(p).Error
	})
}

// GetTopicsByCourseID 查询课程相关的所有帖子
func GetTopicsByCourseID(cid uint) *response.TopicList {
	res := &response.TopicList{}

	return res
}
