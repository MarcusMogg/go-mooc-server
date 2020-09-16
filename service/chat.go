package service

import (
	"server/global"
	"server/model/entity"
	"server/model/response"

	"gorm.io/gorm"
)

// InsertMessage 创建一条新的消息
func InsertMessage(msg *entity.ChatMessage) error {
	return global.GDB.Create(msg).Error
}

// GetUnreadMsgNum 查询所有未读消息,并按照fromID分组
func GetUnreadMsgNum(id uint) []response.UnreadMsgNumResp {
	var res []response.UnreadMsgNumResp
	global.GDB.Model(&entity.ChatMessage{}).Select("from_id, count(*)").
		Group("from_id").Having("status = ? AND to_id = ?", false, id).Scan(&res)
	return res
}

// GetUnreadMsg 查询所有消息,并按照fromID排序
func GetUnreadMsg(id uint) []response.UnreadMsgResp {
	var res []response.UnreadMsgResp
	global.GDB.Transaction(func(tx *gorm.DB) error {
		var tmp []response.ChatMsgResp
		err := global.GDB.Model(&entity.ChatMessage{}).Select("from_id, created_at,msg,m_type").
			Order("from_id").Where("to_id = ?", id).Scan(&tmp).Error
		if err != nil {
			return err
		}

		cur := -1
		for i, j := range tmp {
			if i == 0 || j.FromID != tmp[i-1].FromID {
				result := global.GDB.Model(&entity.ChatMessage{}).
					Where("status = ? AND from_id = ? AND to_id = ?", false, j.FromID, id)
				if result.Error != nil {
					return result.Error
				}
				res = append(res, response.UnreadMsgResp{
					FromID: j.FromID,
					Num:    result.RowsAffected,
				})
				cur++
			}
			res[cur].Msg = append(res[cur].Msg, j)
		}
		return nil
	})
	return res
}

// AckMsg 将 from -> to 之间的消息都设置为已读
func AckMsg(from, to uint) {
	global.GDB.Model(&entity.ChatMessage{}).
		Where("status = ? AND from_id = ? AND to_id = ?", false, from, to).Update("status", true)
}
