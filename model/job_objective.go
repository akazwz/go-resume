package model

// JobObjective 求职意向
type JobObjective struct {
	ResumeID
	JobName  string `json:"job_name" gorm:"comment:求职意向 type:string(30)"`
	BaseCity string `json:"base_city" gorm:"comment:工作城市 type:string(20)"`
	Salary   string `json:"salary" gorm:"comment:期望薪资 type:string(30)"`
}
