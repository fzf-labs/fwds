package errorx

import "net/http"

// @Description: 默认为中文错误码,在NewError()时请传入中文
// @return
var (
	Success = NewError(1, "成功")
	Fail    = NewError(0, "失败")

	//服务级错误
	InternalServerError   = NewError(10001, "服务器发生异常", WithLevel(ErrLevel))
	ErrServiceUnavailable = NewError(10002, "服务不可用", WithLevel(ErrLevel))
	UnknownError          = NewError(10003, "未知错误,请联系管理员", WithLevel(ErrLevel))
	ErrDataException      = NewError(10004, "数据异常", WithLevel(ErrLevel))
)

// 请求相关
var (
	ErrTooManyRequests           = NewError(10101, "请求过多")
	ErrRequestFrequencyIsTooFast = NewError(10105, "请求频率太快了")
	ErrAuthorization             = NewError(10106, "无访问权限")
)

// 参数相关
var (
	ErrBind           = NewError(20001, "参数绑定到结构时发生错误", WithLevel(WarnLevel))
	ErrParam          = NewError(20002, "参数有误", WithLevel(WarnLevel))
	ErrValidation     = NewError(20003, "验证失败")
	ErrNotJsonRequest = NewError(20004, "请使用JSON请求", WithLevel(WarnLevel))
	MissingChannelID  = NewError(20005, "缺少渠道ID")
	MissingDeviceType = NewError(20006, "缺少设备类型")
)

// 数据查询相关
var (
	DataNotExist        = NewError(20100, "数据不存在")
	ErrSql              = NewError(20101, "数据异常", WithLevel(ErrLevel))
	RecordNotFound      = NewError(20102, "记录不存在")
	DuplicateRecords    = NewError(20103, "记录重复")
	NoNeedToMove        = NewError(20104, "无需移动")
	DataFormattingError = NewError(20105, "数据格式化错误", WithLevel(ErrLevel))
	ErrRedis            = NewError(20106, "数据异常!", WithLevel(ErrLevel))
)

// token,签名 ,校验相关
var (
	ErrTokenNotRequest      = NewError(20200, "请求中未携带令牌", WithHttpCode(http.StatusUnauthorized))
	ErrTokenFormat          = NewError(20201, "令牌格式化错误", WithHttpCode(http.StatusUnauthorized))
	ErrToken                = NewError(20202, "错误的token", WithHttpCode(http.StatusUnauthorized))
	ErrTokenInvalid         = NewError(20203, "令牌无效", WithHttpCode(http.StatusUnauthorized))
	ErrTokenExpired         = NewError(20204, "令牌过期", WithHttpCode(http.StatusUnauthorized))
	TokenRefresh            = NewError(20205, "令牌刷新", WithHttpCode(http.StatusUnauthorized))
	ErrTokenRefresh         = NewError(20206, "令牌刷新失败", WithHttpCode(http.StatusUnauthorized))
	TokenVerificationFailed = NewError(20207, "您的登录状态已失效,或在其他设备登录,请您重新登录", WithHttpCode(http.StatusUnauthorized))
	TokenStorageFailed      = NewError(20208, "令牌储存失败", WithHttpCode(http.StatusUnauthorized))
	ErrSignatureParam       = NewError(20209, "签名参数缺失", WithHttpCode(http.StatusUnauthorized))
	WrongTypeOfBusiness     = NewError(20210, "错误的业务类型", WithHttpCode(http.StatusUnauthorized))
)

// 路由权限相关
var (
	NoAccess                            = NewError(20301, "无权限访问")
	UnauthorizedAccess                  = NewError(20302, "非法访问")
	RoutingPermissionVerificationFailed = NewError(20303, "路由权限校验失败")
	MethodNoAccess                      = NewError(20304, "无访问该路由权限")
)

// 文件上传
var (
	FileParsingError            = NewError(20501, "文件解析错误")
	UploadFileDoesNotExist      = NewError(20502, "上传文件不存在")
	FileError                   = NewError(20503, "文件错误")
	FileClassificationException = NewError(20504, "文件分类异常")
	OSSUploadException          = NewError(20505, "OSS上传异常")
)

// 故事相关
var (
	DuplicateMainTitle                 = NewError(20602, "故事标题重复")
	GameNotOpen                        = NewError(20603, "故事未开放")
	YouCanOnlySelectUpTo3TagsForComics = NewError(20604, "故事最多只能选择3个标签")
	GameNeverPlayed                    = NewError(20635, "故事未挑战")
	GameHasPlayed                      = NewError(20636, "故事已重玩,请勿重试")
	GameNeedStartPlay                  = NewError(20637, "请先开始挑战")
	EvaluateLikeRepeated               = NewError(20638, "无需点赞")
)

// 首页
var (
	GameHomepageNotFound   = NewError(20700, "首页不存在")
	GameHomepageDisabled   = NewError(20701, "首页禁用中")
	GameColumnDoesNotExist = NewError(20702, "首页列不存在")
	GameHomepageNotSet     = NewError(20703, "首页未设置")
)

// 超短本
var (
	ShortBookDoNotRepeatTheQuestion = NewError(20800, "请勿重复答题")
	ShortBookDoesNotExist           = NewError(20801, "超短本不存在")
	ShortBookNotNeedLike            = NewError(20802, "无需点赞")
)

// 用户登录注册相关
var (
	SmsSendOverClock                       = NewError(20900, "短信发送超频")
	SmsCodeInvalid                         = NewError(20901, "短信验证码无效")
	SmsCodeExpired                         = NewError(20902, "短信验证码未发送或已失效,请重新发送")
	SmsCodeVerified                        = NewError(20903, "短信验证码已验证")
	SmsRepeatSend                          = NewError(20904, "短信重复发送")
	SmsRequestOverClock                    = NewError(20905, "短信请求超频")
	SmsSendFailed                          = NewError(20906, "短信发送失败")
	SmsTimesLimit                          = NewError(20907, "同一手机号,一天只能发%s次")
	SmsCodeBeenSent                        = NewError(20908, "短信发送频繁，请%s秒后重试")
	EnterTheCorrectPhoneNumber             = NewError(20909, "请填写正确手机号")
	OneClickLoginFailed                    = NewError(20910, "一键登录失败")
	OneClickLoginAuthFailed                = NewError(20911, "一键登录认证失败")
	GuestAccountHasExpired                 = NewError(20912, "游客账号已失效,请重新登录")
	WrongGuestAccount                      = NewError(20913, "错误的游客账号,请重新登录")
	AppleCodeCannotBeEmpty                 = NewError(20914, "苹果登录Code不能为空")
	FailedToAppleUserID                    = NewError(20915, "苹果用户ID获取失败")
	LoginTokenDoesNotExist                 = NewError(20916, "登录token不存在")
	WeChatCodeCannotBeEmpty                = NewError(20917, "微信登录Code不能为空")
	FailedToGetWeChatUserID                = NewError(20918, "微信用户ID获取失败")
	FailedToObtainWeChatUserInformation    = NewError(20919, "微信用户信息获取失败")
	PleaseSignIn                           = NewError(20920, "请登录")
	CaptchaCodeError                       = NewError(20921, "验证码错误")
	QQCodeCannotBeEmpty                    = NewError(20922, "QQ登录Code不能为空")
	FailedToObtainQQChatUserInformation    = NewError(20923, "QQ用户信息获取失败")
	SmsTypeErr                             = NewError(20924, "短信类型错误")
	WeChatMiniProgramUserAcquisitionFailed = NewError(20925, "小程序用户获取失败")
)

// 用户账号
var (
	AccountNotExist                           = NewError(21000, "账号不存在")
	FailedToObtainAccountInformation          = NewError(21001, "账号信息获取失败")
	UserIsLocked                              = NewError(21002, "用户已锁定.请联系客服")
	UserIsLoggedOut                           = NewError(21003, "用户已注销")
	AccountError                              = NewError(21004, "账号错误")
	WrongPassword                             = NewError(21005, "密码错误")
	AccountIsBanned                           = NewError(21006, "账号封禁中")
	TokenGenerationFailed                     = NewError(21007, "Token生成失败")
	UserUpdateFailed                          = NewError(21008, "用户更新失败")
	DuplicateUsername                         = NewError(21009, "用户名重复")
	AbnormalAccountStatus                     = NewError(21010, "账号状态异常,请重试")
	UserBindingException                      = NewError(21011, "用户绑定异常,请重试")
	UserNicknameIsSuspectedOfViolation        = NewError(21012, "用户昵称涉嫌违规")
	UserBindingTypeException                  = NewError(21013, "用户绑定类型异常")
	ThisAccountIsAlreadyBoundByAnotherAccount = NewError(21014, "绑定失败，该账号已被其他账号绑定")
	NoNeedToUnbindTheAccount                  = NewError(21015, "账号无需解绑")
	DuplicateAccountBinding                   = NewError(21016, "账号重复绑定")
	PleaseEnterYourCurrentMobileNumber        = NewError(21017, "请输入您当前的手机号")
	DoNotEnterYourCurrentMobileNumber         = NewError(21018, "请勿输入您当前的手机号")
	OnlyPhoneSignIn                           = NewError(21019, "该账号仅手机号唯一登录方式,不可换绑")
	HasThirdSignIn                            = NewError(21020, "该账号有第三方登录方式,可换绑")
	OnlySignInUnBind                          = NewError(21021, "该社交账号是您登录ALILI的唯一方式,不可解绑")
	IllegalImage                              = NewError(21022, "该头像不符合规范，请修改后重新尝试")
	TimeLimitAvatar                           = NewError(21023, "距离上次修改头像还未达到30天")
	TimeLimitNickName                         = NewError(21024, "距离上次修改昵称还未达到30天")
)

// 通用
var (
	MenuListIsEmpty                = NewError(21100, "菜单列表为空")
	FailedToLoadSensitiveThesaurus = NewError(21101, "敏感词库加载失败")
)

// 用户资产
var (
	AssetException         = NewError(21200, "您的账户资产异常,请联系管理员")
	WrongTypeOfConsumption = NewError(21201, "错误的消费类型")
	DoNotRepeatConsumption = NewError(21202, "请勿重复消费,造成您的资产损失")
	InsufficientBalance    = NewError(21203, "余额不足")
)

// 广告
var (
	IllegalCode = NewError(21500, "非法验证码")
	LimitTime   = NewError(21501, "观看广告次数达到上限")
)

var (
	PayOrderGenerationFailed = NewError(21600, "订单生成失败")
	PayOrderDoesNotExist     = NewError(21601, "订单不存在")
	PayOrderPendingPayment   = NewError(21602, "订单待支付")
	PayOrderClosed           = NewError(21603, "订单已关闭")
	PayOrderPaymentFailed    = NewError(21604, "订单支付失败")
	PayOrderHasBeenRefunded  = NewError(21605, "订单已退款")
	PayOrderException        = NewError(21606, "订单异常")
)
