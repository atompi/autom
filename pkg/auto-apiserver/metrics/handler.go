package metrics

import (
	"github.com/atompi/autom/cmd/auto-apiserver/app/options"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Context struct {
	GinContext *gin.Context
	Options    options.Options
}

type HandlerFunc func(*Context)

func NewHandler(handler HandlerFunc, opts options.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.GinContext = c
		context.Options = opts
		handler(context)
	}
}

func NewPromHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
