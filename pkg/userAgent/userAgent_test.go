package userAgent

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUaSearch(t *testing.T) {
	search := UaSearch("Alili/1.00.067 (com.alili.new; build:10; iOS 15.4.1) Alamofire/5.5.0")
	marshal, err := json.Marshal(search)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
