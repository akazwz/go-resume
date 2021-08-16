package model

import "time"

type Program struct {
	ProgramName string    `json:"program_name" gorm:"comment:项目名称 type:string(30)"`
	Role        string    `json:"role" gorm:"comment:参与角色 type:string(30)"`
	Start       time.Time `json:"start" gorm:"comment:开始时间"`
	End         time.Time `json:"end" gorm:"comment:结束时间"`
	ProgramInfo string    `json:"program_info" gorm:"comment:项目信息 type:longtext"`
}
