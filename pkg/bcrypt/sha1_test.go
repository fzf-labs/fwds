package bcrypt

import (
	"fmt"
	"testing"
)

func TestHmacSha256(t *testing.T) {
	str := "category=tbjc&device_id=123456"
	sha256 := HmacSha256ToHex(str, "abcd")
	fmt.Println(sha256)
}

func TestSHA256(t *testing.T) {
	sha256 := SHA256("category=tbjc&device_id=123456&salt=7c372a90d6e440d8889a852dcc9dbeb5")
	fmt.Println(sha256)
}
