package router

import (
	"github.com/atompi/autom/cmd/auto-apiserver/app/options"
	apisroutergroup "github.com/atompi/autom/pkg/auto-apiserver/apis/router"
	metricsroutergroup "github.com/atompi/autom/pkg/auto-apiserver/metrics/router"
	metrics "github.com/atompi/autom/pkg/metrics/handler"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine, opts options.APIServerOptions) {
	e.Use(metrics.Handler(""))

	rootRouterGroup := e.Group("/")

	metricsroutergroup.MetricsRouter(rootRouterGroup)
	apisroutergroup.ApisRouter(rootRouterGroup, opts)
}
