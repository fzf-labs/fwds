package middleware

import (
	"fwds/pkg/util/uuidutil"

	"github.com/gin-gonic/gin"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuidutil.GenUUID()
		}
		c.Set("X-Request-ID", requestID)
		// Set X-Request-ID header
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}
