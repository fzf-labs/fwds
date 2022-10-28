package signature

import (
	"encoding/json"
	"net/http"
	"time"
)

var _ Signature = (*signature)(nil)

const (
	delimiter = "|"
)

// 合法的 Methods
var methods = map[string]bool{
	http.MethodGet:     true,
	http.MethodPost:    true,
	http.MethodHead:    true,
	http.MethodPut:     true,
	http.MethodPatch:   true,
	http.MethodDelete:  true,
	http.MethodConnect: true,
	http.MethodOptions: true,
	http.MethodTrace:   true,
}

type Signature interface {
	i()
	// Generate 生成签名
	Generate(path string, method string, params json.RawMessage) (sign string, timeStamp string, err error)

	// Verify 验证签名
	Verify(path string, method string, params json.RawMessage, sign string, timeStamp string) (err error)
}

type signature struct {
	key    string
	secret string
	ttl    time.Duration
}

func New(key, secret string, ttl time.Duration) Signature {
	return &signature{
		key:    key,
		secret: secret,
		ttl:    ttl,
	}
}

func (s *signature) i() {}
