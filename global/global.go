package global

import (
	"server/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var MaxVideoNum = 10

var (
	GCONFIG config.Config
	GVP     *viper.Viper
	GDB     *gorm.DB
	UPLOADQUEUE	chan string
)
