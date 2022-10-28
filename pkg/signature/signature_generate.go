package signature

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fwds/pkg/util"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Generate
// path 请求的路径 (不附带 querystring)
func (s *signature) Generate(path string, method string, params json.RawMessage) (sign string, timeStamp string, err error) {
	if err != nil {
		return "", "", err
	}
	if path == "" {
		err = errors.New("path required")
		return
	}

	if method == "" {
		err = errors.New("method required")
		return
	}

	methodName := strings.ToUpper(method)
	if !methods[methodName] {
		err = errors.New("method param error")
		return
	}

	// Date
	timeStamp = strconv.FormatInt(util.Time.NowUnix(), 10)
	// 加密字符串规则
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(path)
	buffer.WriteString(delimiter)
	buffer.WriteString(methodName)
	buffer.WriteString(delimiter)
	buffer.WriteString(string(params))
	buffer.WriteString(delimiter)
	buffer.WriteString(timeStamp)

	// 对数据进行 sha256 加密，并进行 base64 encode
	hash := hmac.New(sha256.New, []byte(s.secret))
	hash.Write(buffer.Bytes())
	sign = base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return
}
