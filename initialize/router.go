package initialize

import (
	"github.com/akazwz/go-resume/api"
	"github.com/akazwz/go-resume/model/response"
	"github.com/akazwz/go-resume/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	router.GET("/i18n", api.I18nTest)
	// Teapot
	router.GET("teapot", response.Teapot)
	routerGroup := router.Group("")
	routers.InitBasicInfoRouter(routerGroup)
	routers.InitResumeRouter(routerGroup)
	return router
}
