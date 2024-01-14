package v1

import (
	"net/http"

	"github.com/atompi/autom/pkg/handler"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *handler.Context) {
	c.GinContext.JSON(http.StatusOK, gin.H{"response": "pong"})
}
