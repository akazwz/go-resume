package model

// Skill 技能特长
type Skill struct {
	ResumeID
	SkillName string `json:"skill_name" gorm:"comment:技能名称 type:string(30)"`
	Level     uint8  `json:"handle_level" gorm:"comment:熟练程度 type:int(1)"`
	SkillInfo string `json:"skill_info" gorm:"comment:技能信息 type:longtext"`
}
