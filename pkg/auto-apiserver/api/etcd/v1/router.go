package v1

import (
	"github.com/atompi/autom/pkg/auto-apiserver/middleware/authorization"

	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup) {
	v1Group := routerGroup.Group("/v1/etcd")
	{
		v1Group.GET("/members", authorization.TokenAuthMiddleware(), Members)
	}
}
