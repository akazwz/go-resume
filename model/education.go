package model

import (
	"time"
)

// Education 教育经历
type Education struct {
	Model
	ResumeID
	School        string    `json:"school" gorm:"comment:学校;type:varchar(30)"`
	Major         string    `json:"major" gorm:"comment:专业;type:varchar(30)"`
	Start         time.Time `json:"start" gorm:"comment:开始就读"`
	Graduation    time.Time `json:"graduation" gorm:"comment:毕业日期"`
	EducationInfo string    `json:"education_info" gorm:"comment:在校信息;type:longtext"`
}
