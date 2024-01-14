package authorization

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Authorization header cannot be empty"})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{"err": "incorrect token format"})
			c.Abort()
			return
		}
		if parts[1] != token {
			c.JSON(http.StatusForbidden, gin.H{"err": "invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
