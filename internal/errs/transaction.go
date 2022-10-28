package errs

import "github.com/pkg/errors"

var (
	OrderDoesNotExist    = errors.New("订单不存在")
	OrderNotifyStatusErr = errors.New("支付通知状态错误")
	PayMethodAppIdNotSet = errors.New("支付配置未设置")
	PayMethodParamErr    = errors.New("支付参数不正确")
)
