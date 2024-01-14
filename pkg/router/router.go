package router

import (
	"github.com/atompi/autom/pkg/options"
	"github.com/gin-gonic/gin"
)

type RouterGroupFunc func(*gin.RouterGroup, options.Options)

func Register(e *gin.Engine, opts options.Options, routerGroups []RouterGroupFunc) {
	rootRouterGroup := e.Group("/")

	for _, routerGroup := range routerGroups {
		routerGroup(rootRouterGroup, opts)
	}
}
