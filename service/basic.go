package service

import (
	"github.com/akazwz/go-resume/global"
	"github.com/akazwz/go-resume/model"
	uuid "github.com/satori/go.uuid"
)

// CreateBasicInfoService create basic info by basic info
func CreateBasicInfoService(bi *model.BasicInfo) (err error, biInter *model.BasicInfo) {
	err = global.GDB.Create(&bi).Error
	return err, bi
}

// DeleteBasicInfoService delete basic info by resume id
func DeleteBasicInfoService(resumeId string) (err error) {
	basicInfo := model.BasicInfo{ResumeID: model.ResumeID{ResumeID: uuid.FromStringOrNil(resumeId)}}
	err = global.GDB.Delete(&basicInfo).Error
	return
}
