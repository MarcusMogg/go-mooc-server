package api

import (
	"fmt"
	"net/http"
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
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// AloneWS 用户单独websocket连接
func AloneWS(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.FailWithMessage("创建websocket连接失败", c)
		return
	}
	var msg request.ChatMsgReq
	err = ws.ReadJSON(&msg)
	if err != nil {
		fmt.Println("chat/alonews:", err)
		ws.Close()
		return
	}
	var user *entity.MUser
	if msg.CMRType == request.USERID {
		user = service.GetUserInfoByID(msg.ToID)
	} else {
		fmt.Println("chat/alonews:没有连接")
		ws.Close()
		return
	}
	// 用户连接时注册到CLIENTS里
	defer func() {
		// 用户退出时删除
		global.CLIENTS.Delete(user.ID)
		ws.Close()
	}()
	global.CLIENTS.Store(user.ID, ws)
	for {
		var msg request.ChatMsgReq
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("chat/alonews:", err)
			break
		}
		if msg.CMRType == request.SENDMSG {
			sendMessage(user.ID, msg.ToID, msg.Msg, entity.MAlone)
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

func sendMessage(from, to uint, msg string, tp entity.MsgType) {
	tows, ok := global.CLIENTS.Load(to)
	msgentity := entity.ChatMessage{
		FromID: from,
		ToID:   to,
		Msg:    msg,
		MType:  tp,
		Status: false,
	}
	service.InsertMessage(&msgentity)
	if ok {
		msgreq := response.ChatMsgResp{
			FromID:   from,
			SendTime: msgentity.CreatedAt,
			Msg:      msg,
			MType:    tp,
		}
		err := (tows.(*websocket.Conn)).WriteJSON(msgreq)
		fmt.Println("chat/sendmessage:", err)
	}
}
