package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-resume/model"
	"go-resume/model/request"
	"go-resume/model/response"
	"go-resume/pkg"
	"go-resume/service"
	"log"
)

// CreateBasicInfo 新建基本信息
// @Summary 新建基本信息
// @Title 新建基本信息
// @Author 赵文卓
// @Description 新建简历基本信息
// @Tags resume basicInfo
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
	if err != nil {
		response.CommonFailed(CodeErrorBindJson, "Bind Json Error", c)
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
		response.CommonFailed(CodeErrorDB, "DB Error", c)
		return
	} else {
		response.Created(CodeSuccessCommon, "Create Success", basicInter, c)
	}
}
