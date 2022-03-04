package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建路由
	r := gin.Default()

	// 2.绑定路由，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello im")
	})

	// 3.监听端口 8088
	r.Run(":8088")
}