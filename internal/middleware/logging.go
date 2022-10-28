package middleware

import (
	"bytes"
	"encoding/json"
	"fwds/pkg/log"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"fwds/internal/response"
	"fwds/pkg/telescope"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		//开始时间
		startTime := time.Now().UTC()

		c.Next()
		//结束时间
		endTime := time.Now().UTC()
		//花费时间
		costSeconds := endTime.Sub(startTime).Seconds()
		//url
		decodedURL, _ := url.QueryUnescape(c.Request.URL.RequestURI())

		//Header 参数 精简
		traceHeader := map[string]string{
			"Content-Type":  c.GetHeader("Content-Type"),
			"User-Agent":    c.GetHeader("User-Agent"),
			"Authorization": c.GetHeader("Authorization"),
			"token":         c.GetHeader("token"),
		}
		rawData, _ := c.GetRawData()
		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))

		t := telescope.GetTelescope(c)
		//请求记录
		t.WithRequest(&telescope.Request{
			TTL:        "un-limit",
			Method:     c.Request.Method,
			DecodedURL: decodedURL,
			Header:     traceHeader,
			Body:       string(rawData),
			ClientIp:   c.ClientIP(),
		})
		//记录返回
		responseBody := bodyLogWriter.body.String()
		var responseCode int
		var responseMsg string
		var responseData interface{}
		if responseBody != "" {
			var jsonResponse response.JsonResponse
			err := json.Unmarshal([]byte(responseBody), &jsonResponse)
			if err == nil {
				responseCode = jsonResponse.Code
				responseMsg = jsonResponse.Msg
				responseData = jsonResponse.Data
			}
		}
		//响应记录
		t.WithResponse(&telescope.Response{
			Header:       c.Writer.Header(),
			HttpCode:     c.Writer.Status(),
			HttpCodeMsg:  http.StatusText(c.Writer.Status()),
			BusinessCode: responseCode,
			BusinessMsg:  responseMsg,
			BusinessData: responseData,
			Body:         responseBody,
			Success:      !c.IsAborted() && c.Writer.Status() == http.StatusOK,
			CostSeconds:  costSeconds,
		})
		log.Logger.Info("log",
			zap.Any("index", t.Index),
			zap.Any("trace_id", t.TraceId),
			zap.Any("request", t.Request),
			zap.Any("response", t.Response),
			zap.Any("sql", t.SQLs),
			zap.Any("redis", t.Redis),
			zap.Any("grpc", t.GRPCs),
			zap.Any("third_party_requests", t.ThirdPartyRequests),
			zap.Any("debugs", t.Debugs),
			zap.Any("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}
