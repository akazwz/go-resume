package model

import (
	uuid "github.com/satori/go.uuid"
)

// Resume 简历
type Resume struct {
	Model
	UserID     uuid.UUID `json:"user_id" gorm:"comment:用户id"`
	ResumeID   uuid.UUID `json:"resume_id" gorm:"comment:简历id;primaryKey;unique"`
	ResumeName string    `json:"resume_name" gorm:"comment:简历名称;type:varchar(30)"`
}

// ResumeID 简历id
type ResumeID struct {
	ResumeID uuid.UUID `json:"resume_id" gorm:"comment:简历id;type:varchar(191);constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ResumeID;references:resume_id"`
}
