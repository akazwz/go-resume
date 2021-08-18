package model

import (
	"gorm.io/gorm"
	"time"
)

// BasicInfo 基本信息
type BasicInfo struct {
	gorm.Model
	ResumeID
	Name            string    `json:"name" gorm:"comment:名字;type:string(30)"`
	Gender          uint8     `json:"gender" gorm:"comment:性别;type:int(1)"`
	Birthday        time.Time `json:"birthday" gorm:"comment:生日"`
	PhoneNumber     string    `json:"phone_number" gorm:"comment:电话号码;type:string(20)"`
	Email           string    `json:"email" gorm:"comment:电子邮箱;type:string(50)"`
	WorkExperience  string    `json:"work_experience" gorm:"comment:工作年限;type:string(20)"`
	ProfilePic      string    `json:"profile_pic" gorm:"comment:照片;type:string(100)"`
	Marriage        uint8     `json:"marriage" gorm:"comment:婚姻状况;type:int(1)"`
	Nation          string    `json:"nation" gorm:"comment:民族;type:string(20)"`
	NativePlace     string    `json:"native_place" gorm:"comment:籍贯;type:string(20)"`
	PoliticalStatus uint8     `json:"political_status" gorm:"comment:政治面貌;type:int(1)"`
	CustomInfo      string    `json:"custom_info" gorm:"comment:自定义信息;type:longtext"`
}
