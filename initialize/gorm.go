package initialize

import (
	"fmt"
	"go-resume/global"
	"go-resume/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitDB() *gorm.DB {
	m := global.CFG.DataBase

	if m.Name == "" {
		return nil
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		m.User,
		m.Password,
		m.Host,
		m.Name,
	)

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return nil
	} else {
		return db
	}
}

func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.BasicInfo{},
	)
	if err != nil {
		os.Exit(0)
	}
}
