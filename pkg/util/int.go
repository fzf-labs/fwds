package util

import "strconv"

var Int = newUtilInt()

type utilInt struct {
}

func newUtilInt() *utilInt {
	return &utilInt{}
}

// Int64ToString 将 int64 转换为 string
func (ut *utilInt) Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// Uint64ToString 将 uint64 转换为 string
func (ut *utilInt) Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func (ut *utilInt) Abs(a int) int {
	return (a ^ a>>31) - a>>31
}
