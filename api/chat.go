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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// AloneWS 用户单独websocket连接
func AloneWS(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.FailWithMessage("创建websocket连接失败", c)
		return
	}
	defer func() {
		// 用户退出时删除
		global.CLIENTS.Delete(user.ID)
		ws.Close()
	}()
	// 用户连接时注册到CLIENTS里
	global.CLIENTS.Store(user.ID, ws)
	for {
		var msg request.ChatMsgReq
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("chat/alonews:", err)
			break
		}
		if msg.CMRType == request.SENDMSG {
			tows, ok := global.CLIENTS.Load(msg.ToID)
			msgentity := entity.ChatMessage{
				FromID: user.ID,
				ToID:   msg.ToID,
				Msg:    msg.Msg,
				MType:  entity.MAlone,
				Status: false,
			}
			service.InsertMessage(&msgentity)
			if ok {
				msgreq := response.ChatMsgResp{
					FromID:   user.ID,
					SendTime: msgentity.CreatedAt,
					Msg:      msg.Msg,
					MType:    entity.MAlone,
				}
				err := (tows.(*websocket.Conn)).WriteJSON(msgreq)
				fmt.Println("chat/alonews:", err)
			}
		} else {
			service.AckMsg(msg.ToID, user.ID)
		}
	}
}

// GetUnreadMsg 获取所有未读信息，按照fromid排序，并把信息设置为已读
func GetUnreadMsg(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	msg := service.GetUnreadMsg(user.ID)
	response.OkWithData(msg, c)
}

// GetUnreadMsgNum 获取所有未读信息的数量，按照fromid分组
func GetUnreadMsgNum(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	msg := service.GetUnreadMsgNum(user.ID)
	response.OkWithData(msg, c)
}
