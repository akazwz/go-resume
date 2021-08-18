package model

import (
	"gorm.io/gorm"
	"time"
)

// WorkExperience 工作经历
type WorkExperience struct {
	gorm.Model
	ResumeID
	Company  string    `json:"company" gorm:"comment:公司;type:varchar(50)"`
	Position string    `json:"position" gorm:"comment:职位;type:varchar(30)"`
	Start    time.Time `json:"start" gorm:"comment:入职时间"`
	End      time.Time `json:"end" gorm:"comment:离职时间"`
	WorkInfo string    `json:"work_info" gorm:"comment:工作信息;type:longtext"`
}
