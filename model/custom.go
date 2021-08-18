package model

import "gorm.io/gorm"

// Custom 自定义模块
type Custom struct {
	gorm.Model
	ResumeID
	Name       string `json:"name" gorm:"comment:自定义名称;type:string(30)"`
	CustomInfo string `json:"custom_info" gorm:"comment:自定义信息;type:longtext"`
}
