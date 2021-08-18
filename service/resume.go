package service

import (
	uuid "github.com/satori/go.uuid"
	"go-resume/global"
	"go-resume/model"
	"time"
)

func CreateResumeService(resume *model.Resume) (err error, r *model.Resume) {
	resume.ResumeID = uuid.NewV4()
	resume.CreatedAt = time.Now()
	err = global.GDB.Create(&resume).Error
	return err, resume
}

func DeleteResumeService(resumeId string) (err error) {
	resume := model.Resume{ResumeID: uuid.FromStringOrNil(resumeId)}
	err = global.GDB.Delete(&resume).Error
	global.GDB.Raw("delete from ")
	return
}
