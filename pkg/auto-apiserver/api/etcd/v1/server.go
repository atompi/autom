package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Members(c *gin.Context) {
	infoStr := "I'm auto-apiserver."
	c.DataFromReader(http.StatusOK, int64(len(infoStr)), gin.MIMEJSON, strings.NewReader(infoStr), nil)
}
