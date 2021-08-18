package model

import (
	"gorm.io/gorm"
	"time"
)

// Education 教育经历
type Education struct {
	gorm.Model
	ResumeID
	School        string    `json:"school" gorm:"comment:学校;type:string(30)"`
	Major         string    `json:"major" gorm:"comment:专业;type:string(30)"`
	Start         time.Time `json:"start" gorm:"comment:开始就读"`
	Graduation    time.Time `json:"graduation" gorm:"comment:毕业日期"`
	EducationInfo string    `json:"education_info" gorm:"comment:在校信息;type:longtext"`
}
