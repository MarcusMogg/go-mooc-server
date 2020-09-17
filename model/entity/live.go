package entity

import (
	"sync"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

// Live 直播信息
type Live struct {
	gorm.Model
	Name         string
	Introduction string
	CourseID     uint
	CourseName   string
	StartTime    string
	EndTime      string
}

// SafeMap 线程安全的map
type SafeMap struct {
	rooms map[uint]map[uint]*websocket.Conn
	mux   sync.Mutex
}

// AddSocket 添加socket
func (c *SafeMap) AddSocket(liveID uint, uid uint, conn *websocket.Conn) {
	c.mux.Lock()
	if c.rooms[liveID] == nil {
		c.rooms[liveID] = map[uint]*websocket.Conn{uid: conn}
		c.rooms[liveID][uid] = conn
	} else {
		c.rooms[liveID][uid] = conn
	}
	c.mux.Unlock()
}

// DropSocket 移除socket
func (c *SafeMap) DropSocket(liveID uint, uid uint) {
	c.mux.Lock()
	delete(c.rooms[liveID], uid)
	c.mux.Unlock()
}

// DropLive 移除直播号
func (c *SafeMap) DropLive(liveID uint) {
	c.mux.Lock()
	delete(c.rooms, liveID)
	c.mux.Unlock()
}
