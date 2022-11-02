package middleware

import (
	"context"
	"fmt"
	"fwds/internal/errno"
	"fwds/internal/response"
	"fwds/pkg/redis"
	"github.com/gin-gonic/gin"
)

func Lock() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetString("uid")
		if uid != "" {
			key := fmt.Sprintf("%s:%s", uid, c.Request.Method)
			lock := redis.NewDefaultLock(key)
			b, err := lock.Lock(c)
			defer func(lock *redis.Lock, ctx context.Context) {
				_ = lock.Unlock(ctx)
			}(lock, c)
			if err != nil {
				response.Json(c, errno.InternalServerError, nil)
				return
			}
			if !b {
				response.Json(c, errno.ErrRequestFrequencyIsTooFast, nil)
			}
		}
		c.Next()
	}
}
