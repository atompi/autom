package router

import (
	"github.com/atompi/autom/pkg/auto-apiserver/handler"
	"github.com/gin-gonic/gin"
)

func MetricsRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("metrics", handler.NewPromHandler())
}
