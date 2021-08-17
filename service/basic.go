package service

import (
	uuid "github.com/satori/go.uuid"
	"go-resume/global"
	"go-resume/model"
	"time"
)

func CreateBasicInfoService(bi *model.BasicInfo) (err error, biInter *model.BasicInfo) {
	bi.ResumeID.CreatedAt = time.Now()
	err = global.GDB.Create(&bi).Error
	return err, bi
}

func CreateResumeService(resume *model.Resume) (err error, r *model.Resume) {
	resume.ResumeID = uuid.NewV4()
	resume.CreatedAt = time.Now()
	err = global.GDB.Create(&resume).Error
	return err, resume
}
