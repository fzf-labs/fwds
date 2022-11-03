package ginutil

import (
	"bytes"
	"encoding/json"
	"fwds/pkg/conv"
	g "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/url"
	"strings"
)

// URI unescape后的uri
func URI(c *g.Context) string {
	uri, _ := url.QueryUnescape(c.Request.URL.RequestURI())
	return uri
}

// RequestInputParams 获取所有参数  json的获取不到
func RequestInputParams(c *g.Context) url.Values {
	// 获取所有参数
	_ = c.Request.ParseForm()
	_ = c.Request.ParseMultipartForm(32 << 20) // 32M
	return c.Request.Form
}

// GetRawData 并重新赋值进去防止下次请求获取不到
func GetRawData(c *g.Context) ([]byte, error) {
	data, err := c.GetRawData()
	if err != nil {
		return nil, err
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data)) // 关键点
	return data, nil
}

// RequestAllToJson 获取所有请求参数并转为 json
func RequestAllToJson(c *g.Context) (json.RawMessage, error) {
	var data json.RawMessage
	var err error
	if c.ContentType() == binding.MIMEJSON {
		data, err = c.GetRawData()
		if err != nil {
			return nil, err
		}
	} else {
		ret := RequestInputParams(c)
		data, err = json.Marshal(ret)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func RequestAllToForm(c *g.Context) (url.Values, error) {
	var ret url.Values
	if c.ContentType() == binding.MIMEJSON {
		jsonMap := make(map[string]interface{})
		err := c.BindJSON(&jsonMap)
		if err != nil {
			return nil, err
		}
		if len(jsonMap) > 0 {
			for k, v := range jsonMap {
				ret.Set(k, conv.String(v))
			}
		}
	} else {
		ret = RequestInputParams(c)
	}
	return ret, nil
}

// RequestAllToJsonWithBody 获取所有请求参数并转为 json 会重新设置原始请求到body中
func RequestAllToJsonWithBody(c *g.Context) (json.RawMessage, error) {
	data, err := c.GetRawData()
	if err != nil {
		return nil, err
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	ret, err := RequestAllToJson(c)
	if err != nil {
		return nil, err
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data)) // 重新设置
	return ret, nil
}

func GetAppKey(c *g.Context) string {
	signatureStr := c.GetHeader("signature")
	if signatureStr == "" {
		return ""
	}
	signatures := strings.Split(signatureStr, " ")
	if len(signatures) != 3 {
		return ""
	}
	return signatures[0]
}
