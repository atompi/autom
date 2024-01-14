package router

import (
	"github.com/atompi/autom/pkg/metrics/handler"
	"github.com/atompi/autom/pkg/options"
	"github.com/gin-gonic/gin"
)

func MetricsRouter(routerGroup *gin.RouterGroup, opts options.Options) {
	routerGroup.GET(opts.APIServer.Metrics.Path, handler.NewPromHandler())
}
