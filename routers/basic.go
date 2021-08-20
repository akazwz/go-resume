package routers

import (
	"github.com/akazwz/go-resume/api"
	"github.com/gin-gonic/gin"
)

func InitBasicInfoRouter(r *gin.RouterGroup) {
	basicInfoRouter := r.Group("/basic-info")
	{
		basicInfoRouter.POST("", api.CreateBasicInfo)
		basicInfoRouter.DELETE("/:id", api.DeleteBasicInfo)
	}
}
