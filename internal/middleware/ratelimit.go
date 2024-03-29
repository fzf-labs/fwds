package middleware

import (
	"fwds/internal/errorx"
	"time"

	"fwds/internal/response"
	"fwds/pkg/ratelimit"

	"github.com/gin-gonic/gin"
)

// Limit
// @Description: 限流器
// @return gin.HandlerFunc
func Limit() gin.HandlerFunc {
	//单机版令牌桶限流
	limiter := ratelimit.NewTokenBucket(10000, time.Second)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			response.Json(c, errorx.ErrTooManyRequests, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
