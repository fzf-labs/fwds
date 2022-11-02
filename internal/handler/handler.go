package handler

import (
	"fwds/pkg/util"
	"fwds/pkg/version"
	nh "net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// getHostname 获取主机名
func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		name = "unknown"
	}
	return name
}

// healthCheckResponse 健康检查响应结构体
type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HealthCheck will return OK if the underlying BoltDB is healthy. At least healthy enough for demoing purposes.
func HealthCheck(c *gin.Context) {
	c.JSON(nh.StatusOK, healthCheckResponse{Status: "UP", Hostname: getHostname()})
}

// RouteNotFound 未找到相关路由
func RouteNotFound(c *gin.Context) {
	c.String(nh.StatusNotFound, "the route not found")
}

// MethodNotFound 方法未找到
func MethodNotFound(c *gin.Context) {
	c.String(nh.StatusNotFound, "method not found")
}

func Home(c *gin.Context) {
	c.JSON(nh.StatusOK, gin.H{
		"status":    "UP",
		"hostname":  getHostname(),
		"client_ip": util.Ip.ClientIP(c.Request),
	})
}

func Version(c *gin.Context) {
	c.JSON(nh.StatusOK, version.Get())
	return
}
