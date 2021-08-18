package model

import (
	"gorm.io/gorm"
	"time"
)

// Internship 实习经历
type Internship struct {
	gorm.Model
	ResumeID
	Company        string    `json:"company" gorm:"comment:公司;type:varchar(30)"`
	Position       string    `json:"position" gorm:"comment:职位;type:varchar(30)"`
	Start          time.Time `json:"start" gorm:"comment:开始时间"`
	End            time.Time `json:"end" gorm:"comment:结束时间"`
	InternshipInfo string    `json:"internship_info" gorm:"comment:实习信息;type:longtext"`
}
