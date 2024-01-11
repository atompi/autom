package metrics

import (
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	engine.GET("metrics", NewPromHandler())
}
