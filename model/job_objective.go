package model

import "gorm.io/gorm"

// JobObjective 求职意向
type JobObjective struct {
	gorm.Model
	ResumeID
	JobName  string `json:"job_name" gorm:"comment:求职意向;type:varchar(30)"`
	BaseCity string `json:"base_city" gorm:"comment:工作城市;type:varchar(20)"`
	Salary   string `json:"salary" gorm:"comment:期望薪资;type:varchar(30)"`
}
