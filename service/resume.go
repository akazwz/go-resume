package service

import (
	"github.com/akazwz/go-resume/global"
	"github.com/akazwz/go-resume/model"
	uuid "github.com/satori/go.uuid"
	"time"
)

// CreateResumeService create resume by resume
func CreateResumeService(resume *model.Resume) (err error, r *model.Resume) {
	resume.ResumeID = uuid.NewV4()
	resume.CreatedAt = time.Now()
	err = global.GDB.Create(&resume).Error
	return err, resume
}

// DeleteResumeService delete resume by resume id
func DeleteResumeService(resumeId string) (err error) {
	resume := model.Resume{ResumeID: uuid.FromStringOrNil(resumeId)}
	err = global.GDB.Delete(&resume).Error
	return
}
