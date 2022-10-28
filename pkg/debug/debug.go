package debug

import (
	"fmt"
	"time"

	"fwds/pkg/color"
	"fwds/pkg/telescope"
	"fwds/pkg/util"

	"github.com/gin-gonic/gin"
)

type Option func(*option)

type Trace = telescope.T

type option struct {
	Telescope *telescope.Telescope
}

func newOption() *option {
	return &option{}
}

func WithContext(ctx *gin.Context) Option {
	return func(opt *option) {
		if ctx != nil {
			opt.Telescope = telescope.GetTelescope(ctx)
		}
	}
}

func Println(key string, value interface{}, options ...Option) {
	ts := time.Now()
	opt := newOption()
	defer func() {
		if opt.Telescope != nil {
			opt.Telescope.AppendDebug(&telescope.Debug{
				Timestamp:   util.Time.NowMicrosecondString(),
				Key:         key,
				Value:       value,
				CostSeconds: time.Since(ts).Seconds(),
			})
		}
	}()
	for _, f := range options {
		f(opt)
	}
	fmt.Println(color.Red(fmt.Sprintf("[Debug] key: %s | value: %v", key, value)))
}

func PrintErr(ctx *gin.Context, err error) {
	getTelescope := telescope.GetTelescope(ctx)
	getTelescope.AppendErr(err)
	fmt.Println(color.Red(fmt.Sprintf("[Err]  %s", err.Error())))
}
