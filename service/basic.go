package service

import (
	"go-resume/global"
	"go-resume/model"
)

func CreateBasicInfoService(bi *model.BasicInfo) (err error, biInter *model.BasicInfo) {
	err = global.GDB.Create(&bi).Error
	return err, bi
}
