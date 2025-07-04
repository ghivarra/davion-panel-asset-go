package corsMiddleware

import (
	"slices"

	"github.com/ghivarra/davion-panel-asset-go/environment"
	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {
	// check origin request
	origin := c.GetHeader("Origin")
	if origin == "" {
		origin = c.GetHeader("origin")
	}

	// check if allowed
	isAllowed := slices.Contains(environment.ALLOWED_HOST, origin)
	if !isAllowed {
		c.AbortWithStatusJSON(403, gin.H{
			"status":  "error",
			"message": "You don't have permission to access this path.",
		})
		return
	}

	// set header
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")

	// if options then send 204 - no content
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	// next
	c.Next()
}
