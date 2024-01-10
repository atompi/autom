package ping

import (
	"net/http"

	"github.com/atompi/autom/pkg/autom-apiserver/api/v1/handler"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *handler.Context) {
	c.GinContext.JSON(http.StatusOK, gin.H{"response": "pong"})
}
