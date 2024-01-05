package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *Context) {
	c.GinContext.JSON(http.StatusOK, gin.H{"response": "pong"})
}
