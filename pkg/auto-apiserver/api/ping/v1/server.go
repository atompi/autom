package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	infoStr := "pong"
	c.DataFromReader(http.StatusOK, int64(len(infoStr)), gin.MIMEJSON, strings.NewReader(infoStr), nil)
}
