package telescope

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"sync"

	"github.com/gin-gonic/gin"
)

const Header = "TRACE-ID"

var _ T = (*Telescope)(nil)

type T interface {
	i()
	ID() string
	WithRequest(req *Request) *Telescope
	WithResponse(resp *Response) *Telescope
	AppendDialog(dialog *Dialog) *Telescope
	AppendSQL(sql *SQL) *Telescope
	AppendRedis(redis *Redis) *Telescope
	AppendGRPC(grpc *Grpc) *Telescope
}

func GetTelescope(ctx *gin.Context) *Telescope {
	t, exists := ctx.Get("Telescope")
	//存在
	if exists {
		//反射一下
		ret, ok := t.(*Telescope)
		if ok {
			return ret
		}
	}
	ret := &Telescope{
		TraceId: ctx.GetString("X-Trace-ID"),
	}
	ret.SetTelescope(ctx)
	return ret
}

// Telescope 记录的参数
type Telescope struct {
	mux                sync.Mutex
	Index              int       `json:"index"`                // 索引ID
	TraceId            string    `json:"trace_id"`             // 链路ID
	Request            *Request  `json:"request"`              // 请求信息
	Response           *Response `json:"response"`             // 返回信息
	ThirdPartyRequests []*Dialog `json:"third_party_requests"` // 调用第三方接口的信息
	Debugs             []*Debug  `json:"debugs"`               // 调试信息
	SQLs               []*SQL    `json:"sqls"`                 // 执行的 SQL 信息
	Redis              []*Redis  `json:"redis"`                // 执行的 Redis 信息
	GRPCs              []*Grpc   `json:"grpc"`                 // 执行的 gRPC 信息
	Err                error
}

// Request 请求信息
type Request struct {
	TTL        string      `json:"ttl"`         // 请求超时时间
	Method     string      `json:"method"`      // 请求方式
	DecodedURL string      `json:"decoded_url"` // 请求地址
	Header     interface{} `json:"header"`      // 请求 Header 信息
	Body       interface{} `json:"body"`        // 请求 Body 信息
	ClientIp   string      `json:"client_ip"`   // 客户端ip
}

// Response 响应信息
type Response struct {
	Header       interface{} `json:"header"`                  // Header 信息
	Body         interface{} `json:"body"`                    // Body 信息
	BusinessCode int         `json:"business_code"`           // 业务码
	BusinessMsg  string      `json:"business_msg"`            // 业务提示信息
	BusinessData interface{} `json:"business_data,omitempty"` // 业务详情
	BusinessErr  interface{} `json:"business_err,omitempty"`  // 业务错误
	HttpCode     int         `json:"http_code"`               // HTTP 状态码
	HttpCodeMsg  string      `json:"http_code_msg"`           // HTTP 状态码信息
	Success      bool        `json:"success"`                 // 请求结果 true or false
	CostSeconds  float64     `json:"cost_seconds"`            // 执行时间(单位秒)
}

func New(id string) *Telescope {
	if id == "" {
		buf := make([]byte, 10)
		_, _ = io.ReadFull(rand.Reader, buf)
		id = hex.EncodeToString(buf)
	}

	return &Telescope{
		TraceId: id,
	}
}

func (t *Telescope) i() {}

// ID 唯一标识符
func (t *Telescope) ID() string {
	return t.TraceId
}

// WithRequest 设置request
func (t *Telescope) WithRequest(req *Request) *Telescope {
	t.Request = req
	return t
}

// WithResponse 设置response
func (t *Telescope) WithResponse(resp *Response) *Telescope {
	t.Response = resp
	return t
}

// AppendDialog 安全的追加内部调用过程dialog
func (t *Telescope) AppendDialog(dialog *Dialog) *Telescope {
	if dialog == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()
	t.Index++
	dialog.Index = t.Index
	t.ThirdPartyRequests = append(t.ThirdPartyRequests, dialog)
	return t
}

// AppendDebug 追加 debug
func (t *Telescope) AppendDebug(debug *Debug) *Telescope {
	if debug == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()
	t.Index++
	debug.Index = t.Index
	t.Debugs = append(t.Debugs, debug)
	return t
}

func (t *Telescope) AppendErr(err error) *Telescope {
	if err == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()
	t.Index++
	t.Err = err
	return t
}
func (t *Telescope) GetErrMsg() string {
	if t.Err == nil {
		return ""
	}
	return t.Err.Error()
}

// AppendSQL 追加 SQL
func (t *Telescope) AppendSQL(sql *SQL) *Telescope {
	if sql == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()
	t.Index++
	sql.Index = t.Index
	t.SQLs = append(t.SQLs, sql)
	return t
}

// AppendRedis 追加 Redis
func (t *Telescope) AppendRedis(redis *Redis) *Telescope {
	if redis == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()
	t.Index++
	redis.Index = t.Index
	t.Redis = append(t.Redis, redis)
	return t
}

// AppendGRPC 追加 gRPC 调用信息
func (t *Telescope) AppendGRPC(grpc *Grpc) *Telescope {
	if grpc == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()
	t.Index++
	grpc.Index = t.Index
	t.GRPCs = append(t.GRPCs, grpc)
	return t
}

func (t *Telescope) SetTelescope(ctx *gin.Context) {
	ctx.Set("Telescope", t)
}
