package auth

import (
	v1 "github.com/anjingGo/im/app/http/controllers/api/v1"
	"github.com/anjingGo/im/pkg/response"
	"github.com/gin-gonic/gin"
)
type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type LoginController struct {
	v1.BaseAPIController
}
func (lgc *LoginController) Register(c *gin.Context)  {
	
}
func (lgc *LoginController) Login(c *gin.Context)  {
	var userLogin Login
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		response.Fail(c, response.CodeServerBusy)
		return
	}

	if userLogin.UserName != "anjing" || userLogin.Password != "admin" {
		response.Fail(c, response.CodeInvalidPassword)
		return
	}
	var userInfo = map[string]interface{}{
		"id" : 1,
		"name": "anjing",
		"token": "tokenstring",
	}
	response.OkWithData(c,userInfo)
}

func (lgc *LoginController) Hi(c *gin.Context)  {
	response.OkWithData(c, "Hi!")
}
