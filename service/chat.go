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
func GetUnreadMsgNum(id uint) int64 {
	var res int64
	global.GDB.Model(&entity.ChatMessage{}).Where("status = ? AND to_id = ?", false, id).Count(&res)
	return res
}

// GetUnreadMsg 查询所有消息,并按照fromID排序
func GetUnreadMsg(id uint) []response.UnreadMsgResp {
	var res []response.UnreadMsgResp
	global.GDB.Transaction(func(tx *gorm.DB) error {
		var tmp []response.ChatMsgResp
		err := tx.Model(&entity.ChatMessage{}).Select("from_id,to_id,created_at,msg,m_type").
			Where("to_id = ? OR from_id = ?", id, id).Scan(&tmp).Error
		if err != nil {
			return err
		}
		var rm map[uint]([]response.ChatMsgResp) = make(map[uint]([]response.ChatMsgResp))
		for _, j := range tmp {
			if j.FromID == id {
				rm[j.ToID] = append(rm[j.ToID], j)
			} else {
				rm[j.FromID] = append(rm[j.FromID], j)
			}
		}
		for i, j := range rm {
			result := global.GDB.Model(&entity.ChatMessage{}).
				Where("status = ? AND from_id = ? AND to_id = ?", false, i, id)
			for k := range j {
				j[k].SendTimeStr = j[k].SendTime.Format(global.TimeTemplateSec)
			}
			res = append(res, response.UnreadMsgResp{
				FromID: i,
				Num:    result.RowsAffected,
				Msg:    j,
			})
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
