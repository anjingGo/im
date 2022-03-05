package main

import (
	"github.com/anjingGo/im/bootstrap"
	"github.com/gin-gonic/gin"
)

type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func main() {
	// 1.创建路由
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	bootstrap.SetupRoute(r)
	// 3.监听端口 8088
	r.Run(":8088")
}