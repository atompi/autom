package router

import (
	"github.com/atompi/autom/cmd/autom-apiserver/app/options"
	"github.com/atompi/autom/pkg/autom-apiserver/api/v1/handler"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine, opts options.Options) {
	engine.GET(opts.API.Path, handler.NewHandler(handler.Handler, opts))
}
