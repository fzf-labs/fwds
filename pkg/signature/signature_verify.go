package signature

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fwds/pkg/conversion"
	"fwds/pkg/util"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func (s *signature) Verify(path string, method string, params json.RawMessage, sign string, timeStamp string) (err error) {
	if path == "" {
		err = errors.New("请求路径不存在")
		return
	}

	if method == "" {
		err = errors.New("请求方法不存在")
		return
	}

	methodName := strings.ToUpper(method)
	if !methods[methodName] {
		err = errors.New("请求方法错误")
		return
	}
	seconds := util.Time.Now().DiffAbsInSeconds(util.Time.CreateFromTimestamp(conversion.Int64(timeStamp)))
	if seconds > int64(s.ttl/time.Second) {
		err = errors.Errorf("接口超时,限时:%v", s.ttl)
		return
	}

	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(path)
	buffer.WriteString(delimiter)
	buffer.WriteString(methodName)
	buffer.WriteString(delimiter)
	buffer.WriteString(string(params))
	buffer.WriteString(delimiter)
	buffer.WriteString(timeStamp)

	// 对数据进行 hmac 加密，并进行 base64 encode
	hash := hmac.New(sha256.New, []byte(s.secret))
	hash.Write(buffer.Bytes())
	digest := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	if sign != digest {
		return errors.New("签名校验不通过")
	}
	return
}
