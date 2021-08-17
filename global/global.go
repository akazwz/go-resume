package global

import (
	"github.com/spf13/viper"
	"go-resume/config"
	"gorm.io/gorm"
)

var (
	CFG config.Conf
	VP  *viper.Viper
	GDB *gorm.DB
)
