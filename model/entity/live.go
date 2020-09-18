package entity

import (
	"sync"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

// Live 直播信息
type Live struct {
	gorm.Model
	Name         string `json:"name"`
	TeacherID    uint   `json:"teacherid"`
	Introduction string `json:"introduction"`
	CourseID     uint   `json:"courseid"`
	CourseName   string `json:"coursename"`
	StartTime    string `json:"starttime"`
	EndTime      string `json:"endtime"`
}

// LiveSign 直播课程签到
type LiveSign struct {
	gorm.Model
	LiveID   uint   `json:"liveid"`
	UserName string `json:"username"`
	UserID   uint   `json:"userid"`
	SignTime string `json:"signtime"`
}

// SafeMap 线程安全的map
type SafeMap struct {
	Rooms map[uint]map[uint]*websocket.Conn
	mux   sync.Mutex
}

// AddSocket 添加socket
func (c *SafeMap) AddSocket(liveID uint, uid uint, conn *websocket.Conn) {
	c.mux.Lock()
	if c.Rooms[liveID] == nil {
		c.Rooms[liveID] = map[uint]*websocket.Conn{uid: conn}
	} else {
		c.Rooms[liveID][uid] = conn
	}
	c.mux.Unlock()
}

// DropSocket 移除socket
func (c *SafeMap) DropSocket(liveID uint, uid uint) {
	c.mux.Lock()
	delete(c.Rooms[liveID], uid)
	c.mux.Unlock()
}

// DropLive 移除直播号
func (c *SafeMap) DropLive(liveID uint) {
	c.mux.Lock()
	delete(c.Rooms, liveID)
	c.mux.Unlock()
}

// Init 初始化函数
func (c *SafeMap) Init() {
	c.Rooms = make(map[uint]map[uint]*websocket.Conn)
}
