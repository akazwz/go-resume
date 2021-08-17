package initialize

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-resume/model/response"
)

func Routers() *gin.Engine {
	router := gin.Default()
	// 路由跨域
	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))
	// 路由不存在的情况
	router.NoRoute(response.NotFound)
	//go-swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Teapot
	router.GET("teapot", response.Teapot)
	return router
}
