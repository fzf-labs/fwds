package middleware

import (
	"fmt"
	"fwds/internal/conf"
	"fwds/pkg/notify"
	"runtime/debug"

	"fwds/internal/errno"
	"fwds/internal/response"
	"fwds/pkg/email"
	"fwds/pkg/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PanicNotify() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stackInfo := string(debug.Stack())
				log.SugaredLogger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", stackInfo))
				dingTalkMsg := fmt.Sprintf("程序:%s发生致命性错误Panic,请及时处理!!!", conf.Conf.App.Name)
				//邮件发送
				subject, body := email.NewNotifyMailData("www.fuzhifei.com", "Panic!", "程序发生致命性错误,请及时处理!!!", "Panic异常告警!!!")
				_ = notify.Send("panic",
					notify.WithDingTalk(dingTalkMsg),
					notify.WithEmail(subject, body),
				)
				response.Json(c, errno.InternalServerError, nil)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
