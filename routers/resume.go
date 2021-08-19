package routers

import (
	"github.com/akazwz/go-resume/api"
	"github.com/gin-gonic/gin"
)

func InitResumeRouter(r *gin.RouterGroup) {
	ResumeRouter := r.Group("resume")
	{
		ResumeRouter.POST("", api.CreateResume)
		ResumeRouter.DELETE("/:id", api.DeleteResume)
	}
}
