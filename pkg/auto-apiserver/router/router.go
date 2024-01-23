package router

import (
	dns "github.com/atompi/autom/pkg/auto-apiserver/apis/dns/v1"
	etcd "github.com/atompi/autom/pkg/auto-apiserver/apis/etcd/v1"
	ping "github.com/atompi/autom/pkg/auto-apiserver/apis/ping/v1"
	metricshandler "github.com/atompi/autom/pkg/metrics/handler"
	metricsroutergroup "github.com/atompi/autom/pkg/metrics/router"
	"github.com/atompi/autom/pkg/options"
	root "github.com/atompi/autom/pkg/router"
	"github.com/gin-gonic/gin"
)

func ApisRouter(routerGroup *gin.RouterGroup, opts options.Options) {
	apisGroup := routerGroup.Group("/apis")

	etcd.Router(apisGroup, opts)
	ping.Router(apisGroup, opts)
	dns.Router(apisGroup, opts)
}

func Register(e *gin.Engine, opts options.Options) {
	routerGroupFuncs := []root.RouterGroupFunc{}

	if opts.APIServer.Metrics.Enable {
		e.Use(metricshandler.Handler(""))
		routerGroupFuncs = append(routerGroupFuncs, metricsroutergroup.MetricsRouter)
	}

	routerGroupFuncs = append(
		routerGroupFuncs,
		ApisRouter,
	)

	root.Register(e, opts, routerGroupFuncs)
}
