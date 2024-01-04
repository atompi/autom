package handler

import (
	"net/http"

	"github.com/atompi/autom/cmd/autom-apiserver/app/options"
	"github.com/gin-gonic/gin"
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

func Handler(c *Context) {
	c.GinContext.JSON(http.StatusOK, gin.H{"status": c.Options})
}
