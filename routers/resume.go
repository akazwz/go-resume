package routers

import (
	"github.com/gin-gonic/gin"
	"go-resume/api"
)

func InitResumeRouter(r *gin.RouterGroup) {
	ResumeRouter := r.Group("resume")
	{
		ResumeRouter.POST("", api.CreateResume)
		ResumeRouter.DELETE("/:id", api.DeleteResume)
	}
}
