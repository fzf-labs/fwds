package requests

// LoginByEmail 邮箱登录
type LoginByEmail struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

// LoginByPhone 手机号登陆传参
type LoginByPhone struct {
	Phone      int `json:"phone" form:"phone" binding:"required,phone"`
	VerifyCode int `json:"password" form:"password" binding:"required"`
}

// EmailRegister 邮箱注册参数
type EmailRegister struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
	Nickname string `json:"nickname" form:"nick_name" binding:"required,gte=3,lte=32"`
}
