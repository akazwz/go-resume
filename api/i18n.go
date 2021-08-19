package api

import (
	"github.com/akazwz/go-i18n/i18n"
	"github.com/akazwz/go-resume/model/response"
	"github.com/gin-gonic/gin"
)

func I18nTest(c *gin.Context) {
	lang := c.Query("lang")
	i := &i18n.I18n{}
	i.SetLang(lang)
	test := i.Trans("test.test1").ToStr()
	response.CommonSuccess(CodeSuccessCommon, test, nil, c)
}
