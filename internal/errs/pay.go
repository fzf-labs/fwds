package errs

import "github.com/pkg/errors"

var (
	OrderIdErr                 = errors.New("应用方订单号错误")
	OrderIdRepeatErr           = errors.New("订单号已存在")
	TransactionIdErr           = errors.New("平台订单号错误")
	PayMethodKeyErr            = errors.New("支付方式错误")
	PayCurrencyErr             = errors.New("支付币种错误")
	PayExpireTimeErr           = errors.New("过期时长错误")
	FailedToGetPaymentDetails  = errors.New("获取支付详细信息失败")
	OrderStatusNoChange        = errors.New("订单状态不能修改")
	OrderStatusErr             = errors.New("订单状态错误")
	RefundAmountExceeded       = errors.New("退款金额超限")
	CannotApplyForARefundTwice = errors.New("不能重复申请退款")
	RefundOrderDoesNotExist    = errors.New("退款订单不存在")
	RefundRepeatNotice         = errors.New("退款重复通知")
)
