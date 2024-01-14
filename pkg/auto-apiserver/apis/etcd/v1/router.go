package v1

import (
	"github.com/atompi/autom/pkg/handler"
	"github.com/atompi/autom/pkg/middleware/authorization"
	"github.com/atompi/autom/pkg/options"

	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup, opts options.Options) {
	EtcdGroup := routerGroup.Group("/etcd/v1")
	{
		EtcdGroup.GET("/members", authorization.TokenAuthMiddleware(opts.APIServer.Token), handler.NewHandler(ListMembersHandler, opts))
	}
}
