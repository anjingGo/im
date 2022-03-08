package bootstrap

import (
	"github.com/anjingGo/im/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine)  {
	routes.RegisterAPIRouter(router)
	routes.RegisterWEBRouter(router)
}