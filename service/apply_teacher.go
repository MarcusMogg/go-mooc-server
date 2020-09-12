package service

import (
	"errors"
	"server/global"
	"server/model/entity"
	"server/model/request"
	"server/model/response"

	"gorm.io/gorm"
)

// InsertApply 添加申请
func InsertApply(a *entity.ApplyTeacher) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("user_id = ? AND state = ?", a.UserID, 0).First(a)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return tx.Create(a).Error
		}
		return errors.New("不要重复申请")
	})
}

// PaginateApply 按页获取内容
func PaginateApply(pagenum, pagesize int) ([]response.ApplyTeacherResp, int64, error) {
	var applys []entity.ApplyTeacher
	var res = make([]response.ApplyTeacherResp, 0, pagesize)
	var total int64
	offset := (pagenum - 1) * pagesize
	result := global.GDB.Model(&entity.ApplyTeacher{}).Count(&total).Offset(offset).Limit(pagesize).Find(&applys)
	if result.Error == nil {
		for _, i := range applys {
			user := GetUserInfoByID(i.UserID)
			res = append(res, response.ApplyTeacherResp{
				ID:       i.ID,
				UserName: user.UserName,
				Email:    user.Email,
				Date:     i.CreatedAt.Format(global.TimeTemplateDay),
				State:    i.State,
			})
		}
	}
	return res, total, result.Error
}

// ChangeApplyState 修改审核状态
func ChangeApplyState(a *request.ApplyAgreeReq) error {
	newState := 1
	if !a.Agree {
		newState = 2
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		var app entity.ApplyTeacher
		err := tx.Where("id = ?", a.ID).First(&app).Error
		if err == nil && app.State == 0 {
			err = tx.Model(&entity.ApplyTeacher{}).Where("id = ?", a.ID).Update("state", newState).Error
			if err != nil {
				return err
			}
			err = tx.Model(&entity.MUser{}).Where("id = ?", app.UserID).Update("role", entity.Teacher).Error
			if err != nil {
				return err
			}
			return nil
		}
		return errors.New("参数错误")
	})
}
