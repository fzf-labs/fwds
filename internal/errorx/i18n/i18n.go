package i18n

import "strconv"

const (
	ZhCN = "zh-CN" // zh-CN 简体中文-中国
	EnUS = "en-US" // en-US 英文-美国
)

var Languages = []string{
	ZhCN,
	EnUS,
}

func GetMessage(code int, lang string) string {
	var msg string
	switch lang {
	case ZhCN:
		msg = zhCNMap[strconv.Itoa(code)]
	case EnUS:
		msg = enUSMap[strconv.Itoa(code)]
	default:
		msg = zhCNMap[strconv.Itoa(code)]
	}
	return msg
}
