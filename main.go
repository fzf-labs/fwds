package main

import (
	"context"
	"fmt"
	"fwds/internal/conf"
	"fwds/internal/job"
	"fwds/internal/router"
	"fwds/pkg/browser"
	"fwds/pkg/color"
	"fwds/pkg/config"
	"fwds/pkg/email"
	"fwds/pkg/log"
	"fwds/pkg/mq"
	"fwds/pkg/trace/jaeger"
	"fwds/pkg/valid"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"github.com/spf13/pflag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func Flag() string {
	var env = pflag.StringP("env", "e", "", "custom setting of environmental parameters")
	pflag.Parse()
	return *env
}

func main() {
	//获取启动时带的参数
	env := Flag()
	//初始化配置
	err := config.Init(env)
	if err != nil {
		panic("config Init err")
	}
	//日志初始化
	log.Init(conf.Conf)
	//数据库初始化
	//db.Init(conf.Conf.Mysql)
	//redis初始化
	//redis.Init(&conf.Conf.Redis)
	//链路跟踪初始化
	jaeger.Init(conf.Conf)
	// Set gin mode.
	gin.SetMode(conf.Conf.App.Mode)
	// 邮箱客户端 初始化
	email.Init(&conf.Conf.Email)
	// 定时任务初始化
	job.Init()
	//nsq 初始化
	mq.Init()
	//
	Profiler()
	// 服务启动信息打印
	ServiceStartupInformationPrinting(conf.Conf)
	// http 服务
	HttpService(conf.Conf)
}

// ServiceStartupInformationPrinting 服务启动信息打印
func ServiceStartupInformationPrinting(cfg *conf.Config) {
	ui := "////////////////////////////////////////////////////////////////////\n//                          _ooOoo_                               //\n//                         o8888888o                              //\n//                         88\" . \"88                              //\n//                         (| ^_^ |)                              //\n//                         O\\  =  /O                              //\n//                      ____/`---'\\____                           //\n//                    .'  \\\\|     |//  `.                         //\n//                   /  \\\\|||  :  |||//  \\                        //\n//                  /  _||||| -:- |||||-  \\                       //\n//                  |   | \\\\\\  -  /// |   |                       //\n//                  | \\_|  ''\\---/''  |   |                       //\n//                  \\  .-\\__  `-`  ___/-. /                       //\n//                ___`. .'  /--.--\\  `. . ___                     //\n//              .\"\" '<  `.___\\_<|>_/___.'  >'\"\".                  //\n//            | | :  `- \\`.;`\\ _ /`;.`/ - ` : | |                 //\n//            \\  \\ `-.   \\_ __\\ /__ _/   .-` /  /                 //\n//      ========`-.____`-.___\\_____/___.-`____.-'========         //\n//                           `=---='                              //\n//      ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^        //\n//                  佛祖保佑       永不宕机     永无BUG              //\n////////////////////////////////////////////////////////////////////"
	fmt.Println(color.Blue(ui))
	fmt.Println(color.Green(fmt.Sprintf("* [register name: %s]", cfg.App.Name)))
	fmt.Println(color.Green(fmt.Sprintf("* [register env: %s]", cfg.App.Mode)))
	fmt.Println(color.Green(fmt.Sprintf("* [register url: %s]", "https://"+cfg.App.Host)))
	intranet := "http://127.0.0.1" + cfg.Http.Addr
	fmt.Println(color.Green(fmt.Sprintf("* [register intranet: %s]", intranet)))
	if cfg.App.Env == "local" {
		_ = browser.Open(intranet)
	}
}

func pprof(cfg *conf.AppConfig) {
	go func() {
		fmt.Printf("Listening and serving PProf HTTP on %s\n", cfg.PprofPort)
		if err := http.ListenAndServe(cfg.PprofPort, http.DefaultServeMux); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatal("listen ListenAndServe for PProf, err: " + err.Error())
		}
	}()
}

func HttpService(cfg *conf.Config) {
	// Create the Gin engine.
	ge := gin.Default()
	//custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("phone", valid.Phone)
	}
	// API Routes.
	router.Load(ge)
	router.Api(ge)
	srv := &http.Server{
		Addr:    cfg.Http.Addr,
		Handler: ge,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatal("listen" + err.Error())
		}
	}()
	gracefulStop(srv)
}

// gracefulStop 优雅退出
// 等待中断信号以超时 5 秒正常关闭服务器
// 官方说明：https://github.com/gin-gonic/gin#graceful-restart-or-stop
func gracefulStop(srv *http.Server) {
	quit := make(chan os.Signal)
	// kill 命令发送信号 syscall.SIGTERM
	// kill -2 命令发送信号 syscall.SIGINT
	// kill -9 命令发送信号 syscall.SIGKILL
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Logger.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Server Shutdown:" + err.Error())
	}
	// 5 秒后捕获 ctx.Done() 信号
	select {
	case <-ctx.Done():
		log.Logger.Info("timeout of 5 seconds.")
	default:
	}
	log.Logger.Info("Server exiting")
}

func Profiler() {
	profiler.Start(profiler.Config{
		ApplicationName: conf.Conf.App.Name,

		// replace this with the address of pyroscope server
		ServerAddress: "http://127.0.0.1:4040",

		// by default all profilers are enabled,
		// but you can select the ones you want to use:
		ProfileTypes: []profiler.ProfileType{
			profiler.ProfileCPU,
			profiler.ProfileAllocObjects,
			profiler.ProfileAllocSpace,
			profiler.ProfileInuseObjects,
			profiler.ProfileInuseSpace,
		},
	})

	// your code goes here
}
