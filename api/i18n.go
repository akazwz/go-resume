package api

import (
	"github.com/gin-gonic/gin"
	"go-resume/model/response"
	"go-resume/pkg/i18n"
)

func I18nTest(c *gin.Context) {
	lang := c.Query("lang")
	if lang == "" {
		lang = "en"
	}
	i := &i18n.I18n{}
	test := i.Trans("test.test1", lang).ToStr()
	response.CommonSuccess(CodeSuccessCommon, test, nil, c)
}
