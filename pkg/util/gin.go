package util

import (
	"bytes"
	"encoding/json"
	"fwds/internal/constants"
	"fwds/pkg/conversion"
	g "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"net/url"
	"strings"
)

var Gin = newGin()

type gin struct {
}

func newGin() *gin {
	return &gin{}
}

// URI unescape后的uri
func (g *gin) URI(c *g.Context) string {
	uri, _ := url.QueryUnescape(c.Request.URL.RequestURI())
	return uri
}

// RequestInputParams 获取所有参数  json的获取不到
func (g *gin) RequestInputParams(c *g.Context) url.Values {
	// 获取所有参数
	_ = c.Request.ParseForm()
	_ = c.Request.ParseMultipartForm(32 << 20) // 32M
	return c.Request.Form
}

//GetRawData 并重新赋值进去防止下次请求获取不到
func (g *gin) GetRawData(c *g.Context) ([]byte, error) {
	data, err := c.GetRawData()
	if err != nil {
		return nil, err
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点
	return data, nil
}

// RequestAllToJson 获取所有请求参数并转为 json
func (g *gin) RequestAllToJson(c *g.Context) (json.RawMessage, error) {
	var data json.RawMessage
	var err error
	if c.ContentType() == binding.MIMEJSON {
		data, err = c.GetRawData()
		if err != nil {
			return nil, err
		}
	} else {
		ret := g.RequestInputParams(c)
		data, err = json.Marshal(ret)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (g *gin) RequestAllToForm(c *g.Context) (url.Values, error) {
	var ret url.Values
	if c.ContentType() == binding.MIMEJSON {
		jsonMap := make(map[string]interface{})
		err := c.BindJSON(&jsonMap)
		if err != nil {
			return nil, err
		}
		if len(jsonMap) > 0 {
			for k, v := range jsonMap {
				ret.Set(k, conversion.String(v))
			}
		}
	} else {
		ret = g.RequestInputParams(c)
	}
	return ret, nil
}

// RequestAllToJsonWithBody 获取所有请求参数并转为 json 会重新设置原始请求到body中
func (g *gin) RequestAllToJsonWithBody(c *g.Context) (json.RawMessage, error) {
	data, err := c.GetRawData()
	if err != nil {
		return nil, err
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	ret, err := g.RequestAllToJson(c)
	if err != nil {
		return nil, err
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 重新设置
	return ret, nil
}

// Lang 获取语言
func (g *gin) Lang(c *g.Context) string {
	language := c.GetHeader("Accept-Language")
	if language != "" {
		return strings.ToLower(language)
	}
	return constants.ZhCN
}

func (g *gin) GetAppKey(c *g.Context) string {
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
