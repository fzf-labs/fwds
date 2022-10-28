package errno

import "net/http"

//
// var
// @Description: 默认为中文错误码,在NewError()时请传入中文
// @return unc
//
var (
	// Common errors
	Success = NewError(0, "成功")
	Fail    = NewError(-1, "请求失败,请联系管理员")

	//服务级错误
	InternalServerError   = NewError(10001, "服务器发生异常")
	ErrServiceUnavailable = NewError(10002, "服务不可用")
	UnknownError          = NewError(10003, "未知错误,请联系管理员")
	ErrDataException      = NewError(10004, "数据异常")

	//请求相关
	ErrTooManyRequests           = NewError(10101, "请求过多")
	ErrAuthorizationNotSet       = NewError(10102, "Authorization not required")
	ErrAuthorizationDateNotSet   = NewError(10103, "Authorization-Date not required")
	ErrAuthorizationFormat       = NewError(10104, "Authorization 格式错误")
	ErrRequestFrequencyIsTooFast = NewError(10105, "请求频率太快了")
	ErrAuthorization             = NewError(10106, "无访问权限")

	//参数校验
	ErrBind           = NewError(20001, "参数绑定到结构时发生错误")
	ErrParam          = NewError(20002, "参数有误")
	ErrValidation     = NewError(20003, "验证失败")
	ErrNotJsonRequest = NewError(20004, "请使用JSON请求")

	//数据库查询
	ErrDatabase = NewError(20101, "数据库错误")

	//token校验
	ErrTokenNotRequest   = NewError(20200, "请求中未携带令牌")
	ErrTokenFormat       = NewError(20201, "令牌格式化错误")
	ErrToken             = NewError(20202, "错误的token")
	ErrTokenInvalid      = NewError(20203, "令牌无效")
	ErrTokenInvalidBlack = NewError(20204, "黑名单中的令牌无效")
	ErrTokenAddBlack     = NewError(20205, "令牌加入黑名单失败")

	//签名相关
	ErrSignatureParam = NewError(20301, "签名参数缺失", WithHttpCode(http.StatusBadRequest))

	//用户校验
	ErrUserNotFound          = NewError(20401, "用户未找到")
	ErrPasswordIncorrect     = NewError(20402, "密码错误")
	ErrEncrypt               = NewError(20403, "密码加密错误")
	ErrAreaCodeEmpty         = NewError(20404, "手机区号不能为空")
	ErrPhoneEmpty            = NewError(20405, "手机号不能为空")
	ErrGenVCode              = NewError(20406, "生成验证码错误")
	ErrVerifyCode            = NewError(20407, "验证码错误")
	ErrEmailOrPassword       = NewError(20408, "邮箱或密码错误")
	ErrTwicePasswordNotMatch = NewError(20409, "两次密码输入不一致")
	ErrRegisterFailed        = NewError(20410, "注册失败")
	ErrRegistered            = NewError(20411, "邮箱已注册")
	ErrCheckFail             = NewError(20412, "邮箱校验失败")
	ErrEmailHasCheck         = NewError(20413, "邮箱已校验")
	ErrEmailForgetPass       = NewError(20414, "忘记密码发送邮箱验证码失败")
	ErrUserEdit              = NewError(20415, "用户修改失败")
	ErrUserArticleList       = NewError(20416, "用户文章列表获取失败")

	//短信
	ErrSendSMS        = NewError(20501, "发送短信错误")
	ErrSendSMSTooMany = NewError(20502, "已超出当日限制，请明天再试")
	//邮件
	ErrSendEmail = NewError(20601, "发送短信错误")

	//文件相关
	ErrBuildName  = NewError(20701, "文件命名错误")
	ErrFile       = NewError(20702, "文件错误")
	ErrFileUpload = NewError(20703, "文件上传失败")
	ErrOssCreate  = NewError(20704, "文件记录创建错误")

	//文章相关
	ErrArticleDetail      = NewError(20801, "文章查询异常")
	ErrArticleStoreInfo   = NewError(20802, "文章创建信息获取失败")
	ErrArticleStore       = NewError(20803, "文章保存失败")
	ErrArticleCommentAdd  = NewError(20804, "文章评论失败")
	ErrArticleCommentLike = NewError(20805, "文章点赞失败")

	//工具相关
	ErrWsConn      = NewError(20901, "websocket连接失败")
	ErrIpLocation  = NewError(20902, "ip地址查询失败")
	ErrSqlToStruct = NewError(20903, "sql解析失败")

	//导航
	ErrWebsiteUrl = NewError(21001, "网址解析错误")
)
