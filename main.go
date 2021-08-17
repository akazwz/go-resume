package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "go-resume/docs"
	"go-resume/global"
	"go-resume/initialize"
	"net/http"
)

// @title Golang Resume Api
// @version 1.0
// @description Golang Resume ASpi
// @termsOfService https://akazwz.com

// @contact.name API Support
// @contact.url https://akazwz.com
// @contact.email akazwz@pm.me

// @license.name MIT
// @license.url https://mit-license.org/

// @host localhost:8000
// @BasePath /
func main() {

	fmt.Println("Hello Gin!")

	// 初始化配置
	global.VP = initialize.InitViper()
	if global.VP == nil {
		fmt.Println("配置文件初始化失败")
	}

	// 设置运行模式
	gin.SetMode(global.CFG.Server.Mode)

	// 初始化数据库
	global.GDB = initialize.InitDB()
	if global.GDB != nil {
		// 新建表
		initialize.CreateTables(global.GDB)
		db, _ := global.GDB.DB()
		// defer关闭数据库
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				fmt.Println("关闭数据库失败")
			}
		}(db)
	}

	// 初始化路由
	routers := initialize.Routers()
	addr := fmt.Sprintf(":%d", global.CFG.Server.Addr)
	server := &http.Server{
		Addr:           addr,
		Handler:        routers,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(`System Serve Start Error`)
	}
}
