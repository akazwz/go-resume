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
	err = global.GDB.Delete(&model.Resume{}, resume).Error
	return
}

// UpdateResumeService update resume by resume
func UpdateResumeService(resume *model.Resume) (err error, rows int64) {
	var r = model.Resume{ResumeID: resume.ResumeID}
	result := global.GDB.Model(&r).
		Updates(model.Resume{ResumeName: resume.ResumeName})
	return result.Error, result.RowsAffected
}

// GetResumesService get all resumes
func GetResumesService() (err error, resumes []model.Resume) {
	err = global.GDB.Find(&resumes).Error
	return err, resumes
}
