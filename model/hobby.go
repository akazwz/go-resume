package model

import "gorm.io/gorm"

// Hobby 兴趣爱好
type Hobby struct {
	gorm.Model
	ResumeID
	Name      string `json:"name" gorm:"comment:爱好名称;type:longtext"`
	HobbyInfo string `json:"hobby_info" gorm:"comment:爱好信息;type:longtext"`
}
