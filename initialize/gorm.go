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
		model.CampusExperience{},
		model.Certificate{},
		model.Custom{},
		model.Education{},
		model.Hobby{},
		model.Internship{},
		model.JobObjective{},
		model.Program{},
		model.Resume{},
		model.SelfAppraisal{},
		model.Skill{},
		model.WorkExperience{},
	)

	if err != nil {
		os.Exit(0)
	}

	tables := make([]string, 0)

	tables = append(tables,
		"basic_infos",
		"campus_experiences",
		"job_objectives",
		"self_appraisals",
		"internships",
		"work_experiences",
		"customs",
		"programs",
		"educations",
		"skills",
		"certificates",
		"hobbies")

	// 添加外键约束
	// 外键约束的名称不能重复
	for i := 0; i < len(tables); i++ {
		tableName := tables[i]
		db.Exec("ALTER TABLE `resume`.`" + tableName + "` " +
			"ADD CONSTRAINT `" + "fk_resume_id_" + tableName + "` FOREIGN KEY (`resume_id`) " +
			"REFERENCES `resume`.`resumes` (`resume_id`) " +
			"ON DELETE CASCADE ON UPDATE CASCADE;")
	}
}
