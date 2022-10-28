package cachekey

import (
	"fwds/pkg/cache"
	"time"
)

var (
	PkgCache        = cache.NewCacheKey("pkg_cache", time.Hour, "pkg 包中的缓存前缀")
	JwtBlack        = cache.NewCacheKey("jwt_black", time.Hour, "jwt 黑名单")
	ForgetPassword  = cache.NewCacheKey("forget_password", time.Hour, "忘记密码")
	SignatureDetail = cache.NewCacheKey("signature_detail", time.Hour, "签名详情")
)
