package routes

import (
	"github.com/anjingGo/im/app/http/controllers/api/v1/auth"
	"github.com/anjingGo/im/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRouter(r *gin.Engine) {
	var v1 *gin.RouterGroup
	v1 = r.Group("/api/v1")
	v1.Use(middleware.JWT())
	{
		authGroup := v1.Group("/auth")
		{
			lgc := new(auth.LoginController)
			authGroup.POST("/login/user-phone",middleware.JWT(),  lgc.Login)
			authGroup.GET("/hi",middleware.JWT(),  lgc.Hi)
		}
	}
}