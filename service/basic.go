package service

import (
	"github.com/akazwz/go-resume/global"
	"github.com/akazwz/go-resume/model"
)

func CreateBasicInfoService(bi *model.BasicInfo) (err error, biInter *model.BasicInfo) {
	err = global.GDB.Create(&bi).Error
	return err, bi
}
