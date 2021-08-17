package main

import (
	"fmt"
	"go-resume/global"
	"go-resume/initialize"
	"net/http"
)

func main() {
	fmt.Println("Hello Gin!")
	global.GDB = initialize.InitDB()
	if global.GDB != nil {
		initialize.CreateTables(global.GDB)
	}

	routers := initialize.Routers()
	addr := ":8000"
	server := &http.Server{
		Addr:           addr,
		Handler:        routers,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(`System Serve Start Error`)
	}
}
