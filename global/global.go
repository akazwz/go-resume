package global

import (
	"go-resume/config"
	"gorm.io/gorm"
)

var (
	CFG config.Conf
	GDB *gorm.DB
)
