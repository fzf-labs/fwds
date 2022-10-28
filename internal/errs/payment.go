package errs

import "github.com/pkg/errors"

var (
	//通用
	ConfigurationError = errors.New("配置错误")

	RefundFailed = errors.New("退款失败")
	//首易信
	PfxErr                                              = errors.New("私钥加载错误")
	CerErr                                              = errors.New("公钥加载错误")
	ParamSortErr                                        = errors.New("参数排序错误")
	SignErr                                             = errors.New("签名错误")
	EncryptErr                                          = errors.New("加密错误")
	DecryptErr                                          = errors.New("解密错误")
	ApiRequestErr                                       = errors.New("接口请求错误")
	HmacNullErr                                         = errors.New("参数签名错误")
	HmacVerifyErr                                       = errors.New("参数签名校验错误")
	UnsupportedPaymentMethodErr                         = errors.New("不支持的支付方式")
	EncryptkeyErr                                       = errors.New("Encryptkey不存在")
	SxyOrderFrequencyTooFastErr                         = errors.New("订单处理频率太快")
	SxyOrderHasBeenFinalizedErr                         = errors.New("订单状态异常,请您重新下单")
	SxyOrderOrderDoesNotExistErr                        = errors.New("订单不存在")
	SxyOrderTimeoutErr                                  = errors.New("订单超时,请您重新下单")
	SxyOrderAlreadyExistsErr                            = errors.New("订单已存在,请您重新下单")
	SxyOrderProcessingFailed                            = errors.New("订单处理异常")
	SxyOrderBankCardDoesNotExist                        = errors.New("首易信银行卡不存在")
	SxyOrderBankCardQueryFailed                         = errors.New("首易信银行卡查询异常")
	SxyOrderInvalidBankCard                             = errors.New("银行卡无效")
	SxyOrderBankCardBindingAbnormal                     = errors.New("银行卡绑定异常")
	SxyOrderSMSVerificationCodeError                    = errors.New("短信验证码错误,验证码验证为一次有效,请您重新发送验证码")
	SxyOrderTheSingleRechargeAmountExceedsTheLimitError = errors.New("单笔充值金额超过限制")
	SxyOrderAlipayChannelAbnormalError                  = errors.New("支付宝通道异常,请联系客服")
	SxyAccessTokenErr                                   = errors.New("access_token err")
	SxyOpenLinkErr                                      = errors.New("open_link err")
)
