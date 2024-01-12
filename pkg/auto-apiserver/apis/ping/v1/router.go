package v1

import (
	"github.com/atompi/autom/cmd/auto-apiserver/app/options"
	"github.com/atompi/autom/pkg/auto-apiserver/handler"
	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup, opts options.APIServerOptions) {
	pingGroup := routerGroup.Group("/ping/v1")
	{
		pingGroup.GET("/ping", handler.NewHandler(PingHandler, opts))
	}
}
