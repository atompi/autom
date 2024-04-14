package router

import (
	dns "github.com/atompi/autom/pkg/auto-apiserver/apis/dns/v1"
	etcd "github.com/atompi/autom/pkg/auto-apiserver/apis/etcd/v1"
	ping "github.com/atompi/autom/pkg/auto-apiserver/apis/ping/v1"
	"github.com/atompi/autom/pkg/options"
	rootrouter "github.com/atompi/autom/pkg/router"
	metricshandler "github.com/atompi/go-kits/metrics/handler"
	"github.com/gin-gonic/gin"
)

func MetricsRouter(routerGroup *gin.RouterGroup, opts options.Options) {
	routerGroup.GET(opts.APIServer.Metrics.Path, metricshandler.NewPromHandler())
}

func ApisRouter(routerGroup *gin.RouterGroup, opts options.Options) {
	apisGroup := routerGroup.Group("/apis")

	etcd.Router(apisGroup, opts)
	ping.Router(apisGroup, opts)
	dns.Router(apisGroup, opts)
}

func Register(e *gin.Engine, opts options.Options) {
	routerGroupFuncs := []rootrouter.RouterGroupFunc{}

	if opts.APIServer.Metrics.Enable {
		e.Use(metricshandler.Handler(""))
		routerGroupFuncs = append(routerGroupFuncs, MetricsRouter)
	}

	routerGroupFuncs = append(
		routerGroupFuncs,
		ApisRouter,
	)

	rootrouter.Register(e, opts, routerGroupFuncs)
}
