package v1

import (
	"github.com/atompi/autom/pkg/auto-apiserver/middleware/authorization"
	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup) {
	v1Group := routerGroup.Group("/v1")
	{
		v1Group.GET("/ping", authorization.TokenAuthMiddleware(), Ping)
	}
}
