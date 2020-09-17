package api

import (
	"fmt"
	"server/global"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
			CourseName:   live.CourseName,
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

// LiveWS 聊天室用户单独websocket连接
func LiveWS(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.FailWithMessage("创建websocket连接失败", c)
		return
	}
	var er request.EnterRoom
	err = ws.ReadJSON(&er)
	if err != nil {
		fmt.Println("live/enter room:", err)
		ws.Close()
		return
	}
	// 用户连接时注册到LIVECLIENTS里
	defer func() {
		// 用户退出时删除
		global.LIVECLIENTS.DropSocket(er.LiveID, er.UID)
		fmt.Println("断开连接")
		ws.Close()
	}()
	global.LIVECLIENTS.AddSocket(er.LiveID, er.UID, ws)
	if er.IsTeacher {
		global.TEACHERS.Store(er.LiveID, er.UID)
	}
	rer := response.EnterRoom{Type: response.ETR, UID: er.UID, UName: er.UName, IsTeacher: er.IsTeacher, Icon: er.Icon}
	for id, tows := range global.LIVECLIENTS.Rooms[er.LiveID] {
		if id != er.UID {
			fmt.Println("转发进入房间的消息", id)
			err := tows.WriteJSON(rer)
			fmt.Println("live/sendmessage:", err)
		}
	}
	fmt.Println("建立连接")
	fmt.Println(global.LIVECLIENTS.Rooms[er.LiveID])
	for {
		var msg request.LiveMsgReq
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("live/empty socket:", err)
			break
		}
		if msg.LiveReqType == request.MSG {
			broadCastChatMsg(er.UID, er.LiveID, msg.ChatData)
		}
	}
}

func broadCastChatMsg(uid, lid uint, msg request.ChatData) {
	for id, tows := range global.LIVECLIENTS.Rooms[lid] {
		if id != uid {
			fmt.Println("转发消息")
			sendChatMsg(tows, msg)
		}
	}
}

func sendChatMsg(tows *websocket.Conn, msg request.ChatData) {
	err := tows.WriteJSON(msg)
	fmt.Println("live/sendmessage:", err)
}
