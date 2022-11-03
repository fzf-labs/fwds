package urlutil

import (
	"fmt"
	"testing"
)

func TestUrlDecode(t *testing.T) {
	str := "username=tizi&password=12345&type=100"
	decode := UrlDecode(str)
	fmt.Println(decode)
}

func TestUrlEncode(t *testing.T) {
	m := map[string]string{"username": "tizi", "password": "12345", "type": "100", "a": "1231"}
	urlEncode := UrlEncode(m)
	fmt.Println(urlEncode)
}
