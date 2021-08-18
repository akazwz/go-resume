package request

type BasicInfo struct {
	ResumeID        string `json:"resume_id" form:"resume_id" biding:"required"`
	Name            string `json:"name" form:"name" binding:"required"`
	Gender          uint8  `json:"gender" form:"gender"` // required 字段不能为空和初始值
	BirthDay        string `json:"birthday" form:"birthday" binding:"required"`
	PhoneNumber     string `json:"phone_number" form:"phone_number" binding:"required"`
	Email           string `json:"email" form:"email" binding:"required"`
	WorkExperience  string `json:"work_experience" form:"work_experience"`
	ProfilePic      string `json:"profile_pic" form:"profile_pic"`
	Marriage        uint8  `json:"marriage" form:"marriage"`
	Nation          string `json:"nation" form:"nation"`
	NativePlace     string `json:"native_place" form:"native_place"`
	PoliticalStatus uint8  `json:"political_status" form:"political_status"`
	CustomInfo      string `json:"custom_info" form:"custom_info"`
}
