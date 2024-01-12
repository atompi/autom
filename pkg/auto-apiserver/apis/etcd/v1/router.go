package v1

import (
	"github.com/atompi/autom/cmd/auto-apiserver/app/options"
	"github.com/atompi/autom/pkg/auto-apiserver/handler"
	"github.com/atompi/autom/pkg/auto-apiserver/middleware/authorization"

	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup, opts options.APIServerOptions) {
	EtcdGroup := routerGroup.Group("/etcd/v1")
	{
		EtcdGroup.GET("/members", authorization.TokenAuthMiddleware(opts), handler.NewHandler(ListMembersHandler, opts))
	}
}
