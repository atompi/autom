package v1

import (
	"github.com/atompi/autom/pkg/handler"
	"github.com/atompi/autom/pkg/middleware/authorization"
	"github.com/atompi/autom/pkg/options"

	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup, opts options.Options) {
	EtcdGroup := routerGroup.Group("/dns/v1")
	{
		EtcdGroup.GET("/origins", authorization.TokenAuthMiddleware(opts.APIServer.Token), handler.NewHandler(ListOriginsHandler, opts))
		EtcdGroup.GET("/records", authorization.TokenAuthMiddleware(opts.APIServer.Token), handler.NewHandler(ListRecordsHandler, opts))
	}
}
