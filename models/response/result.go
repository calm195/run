package response

import (
	"net/http"
	"run/models/constant"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func response(Code int, data interface{}, Msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code: Code,
		Msg:  Msg,
		Data: data,
	})
}

func Ok(c *gin.Context) {
	response(constant.SUCCESS, map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	response(constant.SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	response(constant.SUCCESS, data, "success", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	response(constant.SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	response(constant.ERROR, map[string]interface{}{}, "error", c)
}

func FailWithMessage(message string, c *gin.Context) {
	response(constant.ERROR, map[string]interface{}{}, message, c)
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Result{
		constant.ERROR,
		message,
		nil,
	})
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	response(constant.ERROR, data, message, c)
}
