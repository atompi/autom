package router

import (
	etcd "github.com/atompi/autom/pkg/auto-apiserver/api/etcd/v1"
	ping "github.com/atompi/autom/pkg/auto-apiserver/api/ping/v1"
	"github.com/gin-gonic/gin"
)

type RouterGroup func(*gin.RouterGroup)

func RouterAggregation(routerGroup ...RouterGroup) (routerGroups []RouterGroup) {
	routerGroups = append(routerGroups, routerGroup...)
	return
}

func RouterRegister(routerGroup *gin.RouterGroup, routerGroups []RouterGroup) {
	for _, rg := range routerGroups {
		rg(routerGroup)
	}
}

func ApiRouterGenerator(e *gin.Engine) {
	apiGroup := e.Group("/api")

	routerGroups := RouterAggregation(
		etcd.Router,
		ping.Router,
	)

	RouterRegister(apiGroup, routerGroups)
}
