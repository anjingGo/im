package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code ResCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Result(c *gin.Context, code ResCode, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Message: message,
		Data:    data,
	})
}
func Ok(c *gin.Context) {
	Result(c, CodeSuccess, CodeSuccess.Msg(), nil)
}
func OkWithMessage(c *gin.Context, message string) {
	Result(c, CodeSuccess, message, nil)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, CodeSuccess, CodeSuccess.Msg(), data)
}

func Fail(c *gin.Context, code ResCode) {
	Result(c, code, code.Msg(), nil)
}

func FailWithData(c *gin.Context,code ResCode, data interface{}) {
	Result(c, code, code.Msg(), data)
}
