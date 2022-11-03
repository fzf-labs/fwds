//go:build darwin

package color

import (
	"fmt"
	"math/rand"
	"strconv"
)

var _ = RandomColor()

// RandomColor 生成随机颜色。
func RandomColor() string {
	return fmt.Sprintf("#%s", strconv.FormatInt(int64(rand.Intn(16777216)), 16))
}

// Yellow 黄色
func Yellow(msg string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", msg)
}

// Red 红色
func Red(msg string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", msg)
}

// Redf 红色
func Redf(msg string, arg interface{}) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m %+v\n", msg, arg)
}

// Blue 蓝色
func Blue(msg string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", msg)
}

// Green 绿色
func Green(msg string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", msg)
}

// Greenf 绿色
func Greenf(msg string, arg interface{}) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m %+v\n", msg, arg)
}
