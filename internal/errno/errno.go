package errno

import (
	"fwds/internal/constants"
	"fwds/pkg/errors"
	"net/http"
	"strings"
)

var _ Error = (*err)(nil)

type Error interface {
	// i 为了避免被其他包实现
	i()
	// WithCustomMsg 设置自定义信息 中文语言下 请尽量不使用
	WithCustomMsg(msg string) Error
	// WithErr 设置错误信息
	WithErr(err error) Error
	// GetBusinessCode 获取 Business Code
	GetBusinessCode() int
	// GetHttpCode 获取 HTTP Code
	GetHttpCode() int
	// GetMessage 获取 Msg
	GetMessage(lang string) string
	// GetErr 获取错误信息
	GetErr() error
}

type err struct {
	HttpCode     int    // HTTP Code
	BusinessCode int    // Business Code
	Message      string // 描述信息
	Err          error  // 错误信息
}

type Option func(e *err)

func WithHttpCode(httpCode int) Option {
	return func(e *err) {
		e.HttpCode = httpCode
	}
}
func NewError(businessCode int, msg string, opts ...Option) Error {
	e := &err{
		HttpCode:     http.StatusOK,
		BusinessCode: businessCode,
		Message:      msg,
	}
	if len(opts) > 0 {
		for _, f := range opts {
			f(e)
		}
	}
	return e
}

func (e *err) i() {}

func (e *err) WithErr(err error) Error {
	e.Err = errors.WithStack(err)
	return e
}

func (e *err) WithCustomMsg(msg string) Error {
	e.Message = msg
	return e
}

func (e *err) GetHttpCode() int {
	return e.HttpCode
}

func (e *err) GetBusinessCode() int {
	return e.BusinessCode
}

func (e *err) GetMessage(lang string) string {
	lang = strings.ToLower(lang)
	switch lang {
	case constants.ZhTW:
		zhTw, ok := zhTwMsg[e]
		if ok {
			e.Message = zhTw
		}
	case constants.EnUS:
		enUS, ok := enUSMsg[e]
		if ok {
			e.Message = enUS
		}
	default:

	}
	return e.Message
}

func (e *err) GetErr() error {
	return e.Err
}
