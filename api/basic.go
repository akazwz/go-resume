package api

import (
	"github.com/akazwz/go-i18n/i18n"
	"github.com/akazwz/go-resume/model"
	"github.com/akazwz/go-resume/model/request"
	"github.com/akazwz/go-resume/model/response"
	"github.com/akazwz/go-resume/pkg"
	"github.com/akazwz/go-resume/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
)

// CreateBasicInfo 新建基本信息
// @Summary 新建基本信息
// @Title 新建基本信息
// @Author 赵文卓
// @Description 新建简历基本信息
// @Tags basicInfo
// @Accept json
// @Produce json
// @Param basicInfo body request.BasicInfo true "json"
// @Success 201 {object} model.BasicInfo
// @Failure 400 {object} response.Response
// @Router /basic-info [post]
func CreateBasicInfo(c *gin.Context) {
	var basicInfo request.BasicInfo
	err := c.ShouldBindJSON(&basicInfo)
	log.Println(err)

	lang := c.Query("lang")
	i := &i18n.I18n{}
	i.SetLang(lang)
	bindJsonError := i.Trans("response.bindJsonError").ToStr()
	dbError := i.Trans("response.dbError").ToStr()
	createSuccess := i.Trans("response.createSuccess").ToStr()

	if err != nil {
		response.CommonFailed(CodeErrorBindJson, bindJsonError, c)
		return
	}
	birthDay := pkg.StrToTime(basicInfo.BirthDay, "2006-01-02")

	basic := &model.BasicInfo{
		Name:            basicInfo.Name,
		Gender:          basicInfo.Gender,
		Birthday:        birthDay,
		PhoneNumber:     basicInfo.PhoneNumber,
		Email:           basicInfo.Email,
		WorkExperience:  basicInfo.WorkExperience,
		ProfilePic:      basicInfo.ProfilePic,
		Marriage:        basicInfo.Marriage,
		Nation:          basicInfo.Nation,
		NativePlace:     basicInfo.NativePlace,
		PoliticalStatus: basicInfo.PoliticalStatus,
		CustomInfo:      basicInfo.CustomInfo,
	}

	basic.ResumeID.ResumeID = uuid.FromStringOrNil(basicInfo.ResumeID)

	if err, basicInter := service.CreateBasicInfoService(basic); err != nil {
		response.CommonFailed(CodeErrorDB, dbError, c)
		return
	} else {
		response.Created(CodeSuccessCommon, createSuccess, basicInter, c)
	}
}

// DeleteBasicInfo 删除基本信息
// @Summary 删除基本信息
// @Title 删除基本信息
// @Author 赵文卓
// @Description 删除基本信息
// @Tags basicInfo
// @Param resume_id path int true "Resume ID"
// @Success 204
// @Failure 400 {object} response.Response
// @Router /basic-info/{resume_id} [delete]
func DeleteBasicInfo(c *gin.Context) {
	lang := c.Query("lang")
	i := &i18n.I18n{}
	i.SetLang(lang)
	dbError := i.Trans("response.dbError").ToStr()
	resumeID := c.Param("id")
	err := service.DeleteBasicInfoService(resumeID)
	if err != nil {
		response.CommonFailed(CodeErrorDB, dbError, c)
		return
	}
	response.DeletedNoContent(c)
}
