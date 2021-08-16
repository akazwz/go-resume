package model

import "time"

// Certificate 证书
type Certificate struct {
	ResumeID
	Name            string    `json:"name" gorm:"comment:证书名称 type:string(30)"`
	GetDate         time.Time `json:"get_date" gorm:"comment:取得日期"`
	CertificateInfo string    `json:"certificate_info" gorm:"comment:证书信息 type:longtext"`
}
