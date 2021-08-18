package model

import (
	"time"
)

// BasicInfo 基本信息
type BasicInfo struct {
	Model
	// Resume          Resume    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ResumeID;references:resume_id"`
	ResumeID
	Name            string    `json:"name" gorm:"comment:名字;type:varchar(30)"`
	Gender          uint8     `json:"gender" gorm:"comment:性别;type:tinyint(1)"`
	Birthday        time.Time `json:"birthday" gorm:"comment:生日"`
	PhoneNumber     string    `json:"phone_number" gorm:"comment:电话号码;type:varchar(20)"`
	Email           string    `json:"email" gorm:"comment:电子邮箱;type:varchar(50)"`
	WorkExperience  string    `json:"work_experience" gorm:"comment:工作年限;type:varchar(20)"`
	ProfilePic      string    `json:"profile_pic" gorm:"comment:照片;type:varchar(100)"`
	Marriage        uint8     `json:"marriage" gorm:"comment:婚姻状况;type:tinyint(1)"`
	Nation          string    `json:"nation" gorm:"comment:民族;type:varchar(20)"`
	NativePlace     string    `json:"native_place" gorm:"comment:籍贯;type:varchar(20)"`
	PoliticalStatus uint8     `json:"political_status" gorm:"comment:政治面貌;type:tinyint(1)"`
	CustomInfo      string    `json:"custom_info" gorm:"comment:自定义信息;type:longtext"`
}
