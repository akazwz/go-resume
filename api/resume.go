package api

import (
	"github.com/akazwz/go-i18n/i18n"
	"github.com/akazwz/go-resume/model"
	"github.com/akazwz/go-resume/model/request"
	"github.com/akazwz/go-resume/model/response"
	"github.com/akazwz/go-resume/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
	lang := c.Query("lang")
	i := &i18n.I18n{}
	i.SetLang(lang)
	bindJsonError := i.Trans("response.bindJsonError").ToStr()
	dbError := i.Trans("response.dbError").ToStr()
	createSuccess := i.Trans("response.createSuccess").ToStr()

	var resume request.Resume
	if err := c.ShouldBindJSON(&resume); err != nil {
		response.CommonFailed(CodeErrorBindJson, bindJsonError, c)
		return
	}

	r := model.Resume{
		UserID:     uuid.FromStringOrNil(""),
		ResumeName: resume.ResumeName,
	}

	if err, resumeInter := service.CreateResumeService(&r); err != nil {
		response.CommonFailed(CodeErrorDB, dbError, c)
		return
	} else {
		response.Created(CodeSuccessCommon, createSuccess, resumeInter, c)
	}
}

// DeleteResume 删除简历
// @Summary 删除简历
// @Title 删除简历
// @Author 赵文卓
// @Description 删除简历
// @Tags resume
// @Param resume_id path int true "Resume ID"
// @Success 204
// @Failure 400 {object} response.Response
// @Router /resume/{resume_id} [delete]
func DeleteResume(c *gin.Context) {
	resumeID := c.Param("id")
	lang := c.Query("lang")
	i := &i18n.I18n{}
	i.SetLang(lang)
	dbError := i.Trans("response.dbError").ToStr()

	err := service.DeleteResumeService(resumeID)
	if err != nil {
		response.CommonFailed(CodeErrorDB, dbError, c)
		return
	}
	response.DeletedNoContent(c)
}
