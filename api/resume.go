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
// @Tags resumes
// @Accept json
// @Produce json
// @Param Resume body request.Resume true "json"
// @Param lang query string false "Language"
// @Success 201 {object} model.Resume
// @Failure 400 {object} response.Response
// @Router /resumes [post]
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
// @Tags resumes
// @Param resume_id path string true "Resume ID"
// @Param lang query string false "Language"
// @Success 204
// @Failure 400 {object} response.Response
// @Router /resumes/{resume_id} [delete]
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

// UpdateResume 更新简历
// @Summary 更新简历
// @Title 更新简历
// @Author 赵文卓
// @Description 更新简历
// @Tags resumes
// @Param resume_id path string true "Resume ID"
// @Param Resume body request.Resume true "json"
// @Param lang query string false "Language"
// @Success 201
// @Failure 400 {object} response.Response
// @Router /resumes/{resume_id} [put]
func UpdateResume(c *gin.Context) {
	resumeID := c.Param("id")

	lang := c.Query("lang")
	i := &i18n.I18n{}
	i.SetLang(lang)
	bindJsonError := i.Trans("response.bindJsonError").ToStr()
	dbError := i.Trans("response.dbError").ToStr()
	updateSuccess := i.Trans("response.updateSuccess").ToStr()

	var resume request.Resume
	if err := c.ShouldBindJSON(&resume); err != nil {
		response.CommonFailed(CodeErrorBindJson, bindJsonError, c)
		return
	}

	r := model.Resume{
		ResumeID:   uuid.FromStringOrNil(resumeID),
		ResumeName: resume.ResumeName,
	}

	err, _ := service.UpdateResumeService(&r)
	if err != nil {
		response.CommonFailed(CodeErrorDB, dbError, c)
		return
	}

	response.Created(CodeSuccessCommon, updateSuccess, r, c)
}

// GetResumes 获取所有简历
// @Summary 获取所有简历
// @Title 获取所有简历
// @Author 赵文卓
// @Description 获取所有简历
// @Tags resumes
// @Param lang query string false "Language"
// @Success 200
// @Failure 400 {object} response.Response
// @Router /resumes [get]
func GetResumes(c *gin.Context) {
	lang := c.Query("lang")
	i := &i18n.I18n{}
	i.SetLang(lang)
	dbError := i.Trans("response.dbError").ToStr()
	success := i.Trans("response.success").ToStr()
	err, resumes := service.GetResumesService()
	if err != nil {
		response.CommonFailed(CodeErrorDB, dbError, c)
		return
	}
	response.CommonSuccess(CodeSuccessCommon, success, resumes, c)
}
