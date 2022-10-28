package errs

import "github.com/pkg/errors"

var (
	//支付
	PayNotifyReturnFailed = errors.New("支付通知返回失败")

	//退款
	RefundNotifyReturnFailed = errors.New("退款通知返回失败")
)
