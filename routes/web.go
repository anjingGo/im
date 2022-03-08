package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterWEBRouter(r *gin.Engine) {
	var v1 *gin.RouterGroup
	v1 = r.Group("/")
	v1.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "app.tmpl", gin.H{
			"title": "Main IM WebSite",
		})
	})
}
