package monitor

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pyroscope-io/client/pyroscope"
)

type PyroscopeConfig struct {
	ApplicationName string
	Addr            string
	AuthToken       string
}

func NewPyroscopeServer(cfg *PyroscopeConfig) error {
	if len(cfg.Addr) == 0 {
		return errors.New("pyroscope server not set")
	}
	fmt.Printf("Start pyroscope server, listen addr %s\n", cfg.Addr)
	// 仅当您使用互斥锁或块分析时才需要这两行
	// 请阅读以下说明，了解如何设置这些费率：
	//runtime.SetMutexProfileFraction(5)
	//runtime.SetBlockProfileRate(5)
	_, err := pyroscope.Start(pyroscope.Config{
		//simple.golang.app
		ApplicationName: cfg.ApplicationName,

		// 将其替换为 pyroscope 服务器的地址
		// http://pyroscope-server:4040"
		ServerAddress: cfg.Addr,

		// 您可以通过将其设置为 nil 来禁用日志记录 pyroscope.StandardLogger
		Logger: nil,

		// 可选地，如果启用了身份验证，请指定 API 密钥：
		// AuthToken: os.Getenv("PYROSCOPE_AUTH_TOKEN"),
		AuthToken: cfg.AuthToken,
		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	if err != nil {
		fmt.Printf("Pyroscope start fail:%s", err)
		return err
	}
	fmt.Println("Pyroscope start success")
	return nil
}
