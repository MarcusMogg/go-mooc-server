package global

import (
	"server/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GCONFIG config.Config
	GVP     *viper.Viper
	GDB     *gorm.DB
)
