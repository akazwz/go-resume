package model

import "time"

type WorkExperience struct {
	Company  string    `json:"company" gorm:"comment:公司 type:string(50)"`
	Position string    `json:"position" gorm:"comment:职位 type:string(30)"`
	Start    time.Time `json:"start" gorm:"comment:入职时间"`
	End      time.Time `json:"end" gorm:"comment:离职时间"`
	WorkInfo string    `json:"work_info" gorm:"comment:工作信息 type:longtext"`
}
