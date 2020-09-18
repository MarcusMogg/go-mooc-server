package response

import (
	"sync"
)

// LiveResType 返回类型
type LiveResType int

const (
	// MSG 发送消息
	MSG LiveResType = iota + 1
	// ETR 进入、离开房间
	ETR
	// PTS 推流请求应答
	PTS
	// SPTS 学生请求推流
	SPTS
	// STOP 结束推流
	STOP
)

// EnterRoom 进入直播间返回消息
type EnterRoom struct {
	Type      LiveResType `json:"type"`
	UID       uint        `json:"uid"`
	UName     string      `json:"uname"`
	Icon      string      `json:"icon"`
	IsTeacher bool        `json:"isteacher"`
	IsStudent bool        `json:"isstudent"`
}

// PushStream 推流请求应答
type PushStream struct {
	Type   LiveResType `json:"type"`
	Permit bool        `json:"permit"`
	UID    uint        `json:"uid"`
	UName  string      `json:"uname"`
}

// StudentPushStream 学生请求推流
type StudentPushStream struct {
	Type  LiveResType `json:"type"`
	UID   uint        `json:"uid"`
	UName string      `json:"uname"`
}

// EnterRoomList 消息列表
type EnterRoomList struct {
	Type LiveResType `json:"type"`
	List []EnterRoom
}

// Stop 结束推流
type Stop struct {
	Type LiveResType `json:"type"`
}

// SafeERMap 线程安全的map
type SafeERMap struct {
	Rooms map[uint]map[uint]EnterRoom
	mux   sync.Mutex
}

// AddER 添加EnterRoom
func (c *SafeERMap) AddER(liveID uint, uid uint, ER EnterRoom) {
	c.mux.Lock()
	if c.Rooms[liveID] == nil {
		c.Rooms[liveID] = map[uint]EnterRoom{uid: ER}
	} else {
		c.Rooms[liveID][uid] = ER
	}
	c.mux.Unlock()
}

// DropER 移除ER
func (c *SafeERMap) DropER(liveID uint, uid uint) {
	c.mux.Lock()
	delete(c.Rooms[liveID], uid)
	c.mux.Unlock()
}

// DropLive 移除直播号
func (c *SafeERMap) DropLive(liveID uint) {
	c.mux.Lock()
	delete(c.Rooms, liveID)
	c.mux.Unlock()
}

// Init 初始化函数
func (c *SafeERMap) Init() {
	c.Rooms = make(map[uint]map[uint]EnterRoom)
}
