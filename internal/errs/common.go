package errs

import "github.com/pkg/errors"

var (
	JsonMarshalErr   = errors.New("json参数格式化错误")
	JsonUnmarshalErr = errors.New("json参数解析错误")
)
