package router

import (
	"github.com/atompi/autom/pkg/auto-apiserver/handler"
	"github.com/gin-gonic/gin"
)

func MetricsRouter(routeGroup *gin.RouterGroup) {
	routeGroup.GET("metrics", handler.NewPromHandler())
}
