package signature

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

const (
	key    = "kmq"
	secret = "123456"
	ttl    = time.Minute * 5
)

func TestSignature_Generate(t *testing.T) {
	path := "/v1/pay/method"
	method := "POST"

	params := `{"port":"h5","area":"china","currency":"CNY"}`

	authorization, date, err := New(key, secret, ttl).Generate(path, method, json.RawMessage(params))
	t.Log("signature:", strings.Join([]string{key, date, authorization}, " "))
	t.Log("err:", err)
}

func TestSignature_Verify(t *testing.T) {

	sign := "9BTmNxD/GsDQog3GdLZ9yUfzSnTc9ZJbvILMW+Hsxjo="
	date := "1635008381"

	path := "/v1/pay/method"
	method := "POST"
	params := `{
    "port": "h5",
    "area": "china",
    "currency": "CNY"
}`
	err := New(key, secret, ttl).Verify(path, method, json.RawMessage(params), sign, date)
	t.Log(err)
}
