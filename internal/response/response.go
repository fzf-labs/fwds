package response

import (
	"fwds/internal/errno"
	"fwds/pkg/util/ginutil"

	"github.com/gin-gonic/gin"
)

// JsonResponse api的返回结构体
type JsonResponse struct {
	Code int         `json:"code"`          // 错误码
	Msg  string      `json:"msg"`           // 提示信息
	Data interface{} `json:"data"`          // 返回数据(业务接口定义具体数据结构)
	Err  string      `json:"err,omitempty"` // 只有手动设置才返回
}

// Json json返回
func Json(c *gin.Context, err errno.Error, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	jsonResponse := JsonResponse{
		Code: err.GetBusinessCode(),
		Msg:  err.GetMessage(ginutil.Lang(c)),
		Data: data,
	}
	if err.GetErr() != nil {
		jsonResponse.Err = err.GetErr().Error()
	}
	c.JSON(err.GetHttpCode(), jsonResponse)
}
