package middleware

import (
	"fmt"

	"fwds/internal/errno"
	"fwds/internal/response"
	"fwds/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// JwtMiddleware jwt 鉴权中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取header头中的 Authorization
		authorization := c.Request.Header.Get("Authorization")
		//未获取到token
		if len(authorization) == 0 {
			response.Json(c, errno.ErrTokenNotRequest, nil)
			c.Abort()
			return
		}
		//token解析
		var token string
		_, err := fmt.Scanf(authorization, "Bearer %s", &token)
		if err != nil {
			response.Json(c, errno.ErrTokenFormat, nil)
			c.Abort()
			return
		}
		jc := jwt.NewJC()
		customClaims, err := jc.ParseToken(token)
		if err != nil {
			response.Json(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		//校验是否在黑名单中
		err = jc.CheckBlack(customClaims)
		if err != nil {
			response.Json(c, errno.ErrTokenInvalidBlack, nil)
			c.Abort()
			return
		}
		c.Set("uuid", customClaims.Context.UUID)
		c.Set("nickname", customClaims.Context.NickName)
		c.Next()
	}
}
