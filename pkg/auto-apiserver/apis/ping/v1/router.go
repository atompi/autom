package v1

import (
	"github.com/atompi/autom/pkg/handler"
	"github.com/atompi/autom/pkg/options"
	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup, opts options.Options) {
	pingGroup := routerGroup.Group("/ping/v1")
	{
		pingGroup.GET("/ping", handler.NewHandler(PingHandler, opts))
	}
}
