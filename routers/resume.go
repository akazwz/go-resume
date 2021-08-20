package routers

import (
	"github.com/akazwz/go-resume/api"
	"github.com/gin-gonic/gin"
)

func InitResumeRouter(r *gin.RouterGroup) {
	ResumeRouter := r.Group("resumes")
	{
		ResumeRouter.POST("", api.CreateResume)
		ResumeRouter.DELETE("/:id", api.DeleteResume)
		ResumeRouter.PUT("/:id", api.UpdateResume)
		ResumeRouter.GET("", api.GetResumes)
	}
}
