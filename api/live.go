package api

import (
	"fmt"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"

	"github.com/gin-gonic/gin"
)

// CreateLive 添加直播
func CreateLive(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	var live request.LiveReq
	if err := c.BindJSON(&live); err == nil {
		liveData := entity.Live{
			Name:         live.Name,
			CourseID:     live.CourseID,
			StartTime:    live.StartTime,
			EndTime:      live.EndTime,
			Introduction: live.Introduction,
		}

		if err := service.InsertLive(&liveData); err == nil {
			response.OkWithMessage("创建成功", c)
		} else {
			response.FailWithMessage(fmt.Sprintf("%v", err), c)
		}
	} else {
		response.FailValidate(c)
	}
}

// ReadLive 读取直播信息
func ReadLive(c *gin.Context) {
	var id request.GetByID
	if err := c.BindJSON(&id); err == nil {
		live := service.GetLiveByID(id.ID)
		response.OkWithData(live, c)
	} else {
		response.FailValidate(c)
	}
}

// ReadLiveList 读取课程下的直播列表
func ReadLiveList(c *gin.Context) {
	var id request.GetByID
	if err := c.BindJSON(&id); err == nil {
		lives := service.GetLiveByCourseID(id.ID)
		response.OkWithData(lives, c)
	} else {
		response.FailValidate(c)
	}
}

// GetUserSig 获取直播密钥
/*func GetUserSig(c *gin.Context) {
	var us request.UserSigReq
	if err := c.BindJSON(&us); err == nil {
		utils.GenUserSig()
	}
}
*/
