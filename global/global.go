package global

import (
	"server/config"
	"server/model/entity"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// MaxVideoNum 最大上传视频
var MaxVideoNum = 10

var (
	// GCONFIG 全局配置内容
	GCONFIG config.Config
	// GVP 读取配置
	GVP *viper.Viper
	// GDB 数据库连接
	GDB *gorm.DB
	// UPLOADQUEUE 上传队列
	UPLOADQUEUE chan string
	// CLIENTS 用户套接字
	CLIENTS sync.Map
	// LIVECLIENTS 直播间套接字
	LIVECLIENTS entity.SafeMap
	// TEACHERS 每个直播间的老师
	TEACHERS sync.Map
	// LIVE
)

// TimeTemplateDay 时间转换模板，到天
const TimeTemplateDay = "2006-01-02"

// TimeTemplateSec 时间转换模板，到秒
const TimeTemplateSec = "2006-01-02 15:04:05"
