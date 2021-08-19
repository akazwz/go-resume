package global

import (
	"github.com/akazwz/go-resume/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	CFG config.Conf
	VP  *viper.Viper
	GDB *gorm.DB
)
