package routers

import (
	"github.com/gin-gonic/gin"
	"go-resume/api"
)

func InitBasicInfoRouter(r *gin.RouterGroup) {
	basicInfoRouter := r.Group("/basic-info")
	{
		basicInfoRouter.POST("", api.CreateBasicInfo)
	}
}
