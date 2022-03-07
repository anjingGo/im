package main

import (
	"github.com/anjingGo/im/bootstrap"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func main() {
	// 1.创建路由
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.LoadHTMLGlob("templates/*.tmpl")
	// 如果使用 LoadHTMLFiles 的话这么做（需要列举所有需要加载的文件，不如上述 LoadHTMLGlob 模式匹配方便）：
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	bootstrap.SetupRoute(r)
	// 3.监听端口 8088
	r.Run(":8088")
}