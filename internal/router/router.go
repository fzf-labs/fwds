package router

import (
	"fwds/internal/conf"
	"fwds/internal/constants"
	"fwds/internal/handler"
	"fwds/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	//全局中间件
	//g.Use(middleware.PanicNotify())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.Cors())        //跨域
	g.Use(middleware.Logging())     //日志
	g.Use(middleware.RequestID())   //请求ID记录
	g.Use(middleware.Trace())       //链路跟踪
	g.Use(middleware.PanicNotify()) //panic
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(handler.RouteNotFound)
	g.NoMethod(handler.MethodNotFound)

	// 静态资源，主要是图片
	g.Static("/static", "./static")

	// swagger http docs
	if conf.Conf.App.Env == constants.EnvLocal || conf.Conf.App.Env == constants.EnvDev {
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// pprof router 性能分析路由
	// 默认关闭，开发环境下可以打开
	// 访问方式: HOST/debug/pprof
	// 通过 HOST/debug/pprof/profile 生成profile
	// 查看分析图 go tool pprof -http=:5000 profile
	// pprof.Register(g)

	// HealthCheck 健康检查路由
	g.GET("/health", handler.HealthCheck)
	// metrics router 可以在 prometheus 中进行监控
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//根路由
	g.GET("/", handler.Home)
	//代码版本
	g.GET("/version", handler.Version)
	return g
}
