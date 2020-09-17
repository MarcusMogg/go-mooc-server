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
		tx.Model(&entity.Post{}).Where("id = ?", p.TopicID).Update("num", gorm.Expr("num + ?", 1))
		return tx.Create(p).Error
	})
}

// GetTopicsByCourseID 查询课程相关的所有帖子
func GetTopicsByCourseID(pagenum, pagesize int, cid uint) *response.TopicList {
	res := &response.TopicList{}
	offset := (pagenum - 1) * pagesize
	var tot int64
	global.GDB.Model(&entity.Topic{}).Where("cid = ?", cid).Count(&tot).Offset(offset).Limit(pagesize).Find(&res.Topics)
	res.Num = uint(tot)
	return res
}

// GetTopicDetail 查询主题详情
func GetTopicDetail(pagenum, pagesize int, id uint) *response.TopicDetailResp {
	res := &response.TopicDetailResp{}
	offset := (pagenum - 1) * pagesize
	var tot int64
	global.GDB.Model(&entity.Post{}).Where("topic_id = ?", id).Count(&tot).Offset(offset).Limit(pagesize).Find(&res.Posts)
	res.Num = uint(tot)
	var t entity.Topic
	global.GDB.Where("id = ?", id).First(&t)
	res.Title = t.Title
	res.CourseIDReq.ID = t.CID
	return res
}

// DropTopic 删除主题
func DropTopic(id uint) {
	global.GDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", id).Delete(&entity.Topic{}).Error
		if err != nil {
			return nil
		}
		return tx.Where("topic_id = ?", id).Delete(&entity.Post{}).Error
	})
}

// DropPost 删除帖子
func DropPost(id uint) {
	global.GDB.Transaction(func(tx *gorm.DB) error {
		var topicid uint
		tx.Model(&entity.Post{}).Select("topic_id").Where("id = ?", id).Scan(&topicid)
		tx.Model(&entity.Post{}).Where("id = ?", topicid).Update("num", gorm.Expr("num - ?", 1))
		return tx.Where("id = ?", id).Delete(&entity.Post{}).Error
	})
}

// SetTopicTop 设置置顶状态
func SetTopicTop(id uint, status bool) {
	global.GDB.Model(&entity.Topic{}).Where("id = ?", id).Update("top", status)
}

// SetTopicImport 设置精华状态
func SetTopicImport(id uint, status bool) {
	global.GDB.Model(&entity.Topic{}).Where("id = ?", id).Update("important", status)
}

// Like 点赞
func Like(uid, pid uint) {
	global.GDB.Transaction(func(tx *gorm.DB) error {
		l := entity.UserLike{
			UID: uid,
			PID: pid,
		}
		tx.Model(&entity.Post{}).Where("id = ?", pid).Update("num", gorm.Expr("num + ?", 1))
		return tx.Create(l).Error
	})
}

// UnLike 点赞
func UnLike(uid, pid uint) {
	global.GDB.Transaction(func(tx *gorm.DB) error {
		l := entity.UserLike{
			UID: uid,
			PID: pid,
		}
		tx.Model(&entity.Post{}).Where("id = ?", pid).Update("num", gorm.Expr("num - ?", 1))
		return tx.Delete(l).Error
	})
}
