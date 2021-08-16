package model

// SelfAppraisal 自我评价
type SelfAppraisal struct {
	ResumeID
	KeyWord string `json:"key_word" gorm:"comment:关键词 type:longtext"`
	Info    string `json:"info" gorm:"comment:自我评价 type:longtext"`
}
