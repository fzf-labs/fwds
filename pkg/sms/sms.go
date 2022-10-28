package sms

import (
	"fwds/internal/conf"
)

type Sms interface {
	GetBusiness(key string) Sms
	Send(phone []string, templateParam interface{}) error
}

func NewDefaultSms() Sms {
	switch conf.Conf.Sms.Use {
	case "AliYun":
		return NewAliYun()
	case "Tencent":
		return NewTencentCloud()
	default:
		panic("sms driver err")
	}
}
