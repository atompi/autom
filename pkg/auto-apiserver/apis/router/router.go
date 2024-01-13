package router

import (
	"github.com/atompi/autom/cmd/auto-apiserver/app/options"
	etcd "github.com/atompi/autom/pkg/auto-apiserver/apis/etcd/v1"
	ping "github.com/atompi/autom/pkg/auto-apiserver/apis/ping/v1"
	"github.com/gin-gonic/gin"
)

func ApisRouter(routerGroup *gin.RouterGroup, opts options.APIServerOptions) {
	apisGroup := routerGroup.Group("/apis")

	etcd.Router(apisGroup, opts)
	ping.Router(apisGroup, opts)
}
