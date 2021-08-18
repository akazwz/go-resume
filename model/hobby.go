package model

// Hobby 兴趣爱好
type Hobby struct {
	Model
	ResumeID
	Name      string `json:"name" gorm:"comment:爱好名称"`
	HobbyInfo string `json:"hobby_info" gorm:"comment:爱好信息;type:longtext"`
}
