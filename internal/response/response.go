package response

import (
	"fwds/internal/errorx"
	"fwds/internal/errorx/i18n"
	"fwds/pkg/util/validutil"
	"github.com/pkg/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Code   int         `json:"code"`             // 错误码
	Msg    string      `json:"msg"`              // 描述信息
	Data   interface{} `json:"data"`             // 返回信息
	Err    string      `json:"err,omitempty"`    // 错误信息
	Detail string      `json:"detail,omitempty"` // 错误堆栈
}

func HttpSuccess(resp interface{}, lang string) *HttpResponse {
	r := &HttpResponse{
		Code: errorx.Success.GetBusinessCode(),
		Msg:  errorx.Success.GetMessage(lang),
		Data: resp,
	}
	if validutil.IsZero(r.Data) {
		r.Data = gin.H{}
	}
	return r
}

func httpError(err *errorx.BusinessErr, lang string) *HttpResponse {
	r := &HttpResponse{
		Code:   err.GetBusinessCode(),
		Msg:    err.GetMessage(lang),
		Data:   err.GetErrData(),
		Err:    err.GetErrMsg(),
		Detail: err.GetDetail(),
	}
	return r
}

// Lang 获取语言
func Lang(c *gin.Context) string {
	language := c.GetHeader("Accept-Language")
	if language != "" {
		return strings.ToLower(language)
	}
	return i18n.ZhCN
}

// Json json返回
func Json(c *gin.Context, err error, data interface{}) {
	if err == nil {
		c.JSON(http.StatusOK, HttpSuccess(data, Lang(c)))
	} else {
		var e *errorx.BusinessErr
		causeErr := errors.Cause(err)
		//自定义错误类型检测
		if businessErr, ok := causeErr.(*errorx.BusinessErr); ok {
			e = businessErr
		} else {
			e = errorx.InternalServerError.WithDetail(err)
		}
		c.JSON(e.GetHttpCode(), httpError(e, Lang(c)))
	}
}
