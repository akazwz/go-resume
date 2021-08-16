package model

import "time"

type CampusExperience struct {
	ExperienceName string    `json:"experience_name" gorm:"comment:经历名称 type:string(30)"`
	Role           string    `json:"role" gorm:"comment:角色 type:string(30)"`
	Start          time.Time `json:"start" gorm:"comment:开始时间"`
	End            time.Time `json:"end" gorm:"comment:结束时间"`
	ExperienceInfo string    `json:"experience_info" gorm:"comment:校园经历信息 type:longtext"`
}
