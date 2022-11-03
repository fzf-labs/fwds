package constants

import (
	"fwds/pkg/cache"
	"time"
)

// 缓存key前缀
var (
	UUID      = cache.NewCacheKey("uuid", time.Hour, "uuid")
	DL        = cache.NewCacheKey("dl", time.Second*5, "分布式锁")
	JwtBlack  = cache.NewCacheKey("jwt_black", time.Hour, "jwt 黑名单")
	Sms       = cache.NewCacheKey("sms", time.Minute*5, "短信验证")
	SmsDayNum = cache.NewCacheKey("sms_day_num", time.Minute*5, "短信发送次数")
	FileUrl   = cache.NewCacheKey("file_url", time.Hour*24, "文件的url")
)
