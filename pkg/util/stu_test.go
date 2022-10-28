package util

import (
	"fmt"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestStu_StructToTagValue(t *testing.T) {
	user := User{
		Name: "fzf",
		Age:  123,
	}
	value := Stu.StructToJsonTagValue(user)
	fmt.Println(value)
}
