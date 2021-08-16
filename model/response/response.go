package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"` //omitempty nil or default
	Msg  string      `json:"msg,omitempty"`
}

const (
	SUCCESS  = 2000
	ERROR    = 4000
	PROGRESS = 2020
)

// NotFound 路由不存在
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		Code: ERROR,
		Msg:  "404 not found",
	})
}
