package global

import (
	"server/config"
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
)

// TimeTemplateDay 时间转换模板，到天
const TimeTemplateDay = "2006-01-02"

// TimeTemplateSec 时间转换模板，到秒
const TimeTemplateSec = "2006-01-02 15:04:05"
