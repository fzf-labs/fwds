package router

import (
	_ "fwds/docs" // swagger must import it.
	"fwds/internal/api/http/v1/test"
	"fwds/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Api(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(mw...)
	g.GET("/test", test.HelloWorld)
	u := g.Group("/v1")
	//jwt
	u.Use(middleware.JwtMiddleware())
	{
	}
	return g
}
