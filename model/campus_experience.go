package model

import (
	"gorm.io/gorm"
	"time"
)

// CampusExperience 校园经历
type CampusExperience struct {
	gorm.Model
	ResumeID
	ExperienceName string    `json:"experience_name" gorm:"comment:经历名称;type:varchar(30)"`
	Role           string    `json:"role" gorm:"comment:角色;type:varchar(30)"`
	Start          time.Time `json:"start" gorm:"comment:开始时间"`
	End            time.Time `json:"end" gorm:"comment:结束时间"`
	ExperienceInfo string    `json:"experience_info" gorm:"comment:校园经历信息;type:longtext"`
}
