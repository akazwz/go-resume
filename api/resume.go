package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-resume/model"
	"go-resume/model/request"
	"go-resume/model/response"
	"go-resume/service"
)

// CreateResume 新建简历
// @Summary 新建简历
// @Title 新建简历
// @Author 赵文卓
// @Description 新建简历
// @Tags resume
// @Accept json
// @Produce json
// @Param Resume body request.Resume true "json"
// @Success 201 {object} model.Resume
// @Failure 400 {object} response.Response
// @Router /resume [post]
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