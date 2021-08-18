package service

import (
	"go-resume/global"
	"go-resume/model"
	"time"
)

func CreateBasicInfoService(bi *model.BasicInfo) (err error, biInter *model.BasicInfo) {
	bi.ResumeID.CreatedAt = time.Now()
	err = global.GDB.Create(&bi).Error
	return err, bi
}
