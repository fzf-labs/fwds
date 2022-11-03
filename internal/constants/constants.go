package constants

const (
	// DevMode means development mode.
	DevMode = "dev"
	// TestMode means test mode.
	TestMode = "test"
	// PreMode means pre-release mode.
	PreMode = "pre"
	// ProMode means production mode.
	ProMode = "pro"
)

// 常量定义
const (
	HeaderAcceptLanguage = "Accept-Language" // 语言标识
	HeaderXDeviceType    = "X-Device-Type"   // 设备类型（Http头部字段）
	HeaderXDeviceID      = "X-Device-Id"     // 设备ID（Http头部字段）
	HeaderXRealIP        = "X-Real-IP"       // 客户端IP（Http头部字段）
	HeaderXAppVersion    = "X-App-Version"   // App版本号（Http头部字段）

	IdentityTypeWeChat = 1 //用户第三方登录-微信
	IdentityTypeApple  = 2 //用户第三方登录-苹果
	IdentityTypeQQ     = 3 //用户第三方登录-qq

	//jwt业务类型
	JwtTypeApp   = "App"
	JwtTypeAdmin = "Admin"

	//短信发送类型
	SmsTypeRegister     = "register"
	SmsTypeAccountCheck = "account_check"
	SmsTypeAccountBind  = "account_bind"

	Ios     = 0 // 用户设备类型（IOS） ios
	Android = 1 // 用户设备类型（安卓） android
)
