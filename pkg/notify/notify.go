package notify

import (
	"fwds/internal/conf"
	"fwds/pkg/email"
	"fwds/pkg/log"
	"fwds/pkg/sms"
	"fwds/pkg/webhook/dingtalk"
	"github.com/pkg/errors"
	"strings"
)

type option struct {
	Sms         interface{}
	Subject     string
	Body        string
	dingTalkMsg string
}

type Option func(msg *option)

func Send(key string, o ...Option) error {
	business, ok := conf.Conf.Notify[key]
	if !ok {
		return errors.New("notify key not found")
	}
	opt := new(option)
	if len(o) > 0 {
		for _, f := range o {
			f(opt)
		}
	}
	if business.Email != "" && opt.Subject != "" && opt.Body != "" {
		emails := strings.Split(business.Email, ",")
		for _, s := range emails {
			s := s
			go func() {
				err := email.Send(s, opt.Subject, opt.Body)
				if err != nil {
					log.SugaredLogger.Errorf("notify email send err %v", err)
				}
			}()
		}

	}
	if business.Sms != "" && opt.Sms != nil {
		smsList := strings.Split(business.Sms, ",")
		err := sms.NewDefaultSms().GetBusiness("err").Send(smsList, opt.Sms)
		if err != nil {
			log.SugaredLogger.Errorf("notify sms send err %v", err)
			return err
		}
	}
	if business.DingTalk != "" && opt.dingTalkMsg != "" {
		err := dingtalk.SendText(business.DingTalk, opt.dingTalkMsg)
		if err != nil {
			log.SugaredLogger.Errorf("notify dingtalk send err %v", err)
			return err
		}
	}
	return nil
}

func WithEmail(subject, body string) Option {
	return func(o *option) {
		o.Subject = subject
		o.Body = body
	}
}

func WithSms(sms interface{}) Option {
	return func(o *option) {
		o.Sms = sms
	}
}

func WithDingTalk(msg string) Option {
	return func(o *option) {
		o.dingTalkMsg = msg
	}
}
