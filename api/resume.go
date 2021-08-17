package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-resume/model"
	"go-resume/model/request"
	"go-resume/model/response"
	"go-resume/service"
)

func CreateResume(c *gin.Context) {
	var resume request.Resume
	if err := c.ShouldBindJSON(&resume); err != nil {
		response.CommonFailed(CodeErrorBindJson, "Bind Json Error", c)
		return
	}

	r := model.Resume{
		UserID:     uuid.FromStringOrNil(""),
		ResumeName: "",
	}

	if err, resumeInter := service.CreateResumeService(&r); err != nil {
		response.CommonFailed(CodeErrorDB, "DB Error", c)
		return
	} else {
		response.Created(CodeSuccessCommon, "Success", resumeInter, c)
	}
}
