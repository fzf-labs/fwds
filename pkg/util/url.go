package util

import (
	"strings"

	"github.com/qiniu/go-sdk/v7/storage"
)

// URL
// var Url = newUrl
// @Description: url工具
//
var URL = newUrl()

type urlUtil struct {
}

func newUrl() *urlUtil {
	return &urlUtil{}
}
func (uu *urlUtil) CheckUrl(url string) bool {
	return Reg.RegexpUrlFormat(url)
}

// GetDefaultAvatarURL 获取默认头像
func (uu *urlUtil) GetDefaultAvatarURL(cdnURL string) string {
	uri := "/default/avatar.jpg"
	return uu.GetQiNiuPublicAccessURL(cdnURL, uri)
}

// GetAvatarURL user's avatar, if empty, use default avatar
func (uu *urlUtil) GetAvatarURL(cdnURL, key string) string {
	if key == "" {
		return uu.GetDefaultAvatarURL(cdnURL)
	}
	if strings.HasPrefix(key, "https://") {
		return key
	}
	return uu.GetQiNiuPublicAccessURL(cdnURL, key)
}

// GetQiNiuPublicAccessURL 获取七牛资源的公有链接
// 无需配置bucket, 域名会自动到域名所绑定的bucket去查找
func (uu *urlUtil) GetQiNiuPublicAccessURL(cdnURL, path string) string {
	domain := cdnURL
	key := strings.TrimPrefix(path, "/")

	publicAccessURL := storage.MakePublicURL(domain, key)

	return publicAccessURL
}
