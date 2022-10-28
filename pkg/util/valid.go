package util

import (
	"reflect"
)

var Valid = newValid()

type valid struct {
}

func newValid() *valid {
	return &valid{}
}

// IsZero 检查是否是零值
func (uv *valid) IsZero(i ...interface{}) bool {
	for _, j := range i {
		v := reflect.ValueOf(j)
		if isZero(v) {
			return true
		}
	}
	return false
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return v.IsNil()
	case reflect.Invalid:
		return true
	default:
		z := reflect.Zero(v.Type())
		return reflect.DeepEqual(z.Interface(), v.Interface())
	}
}
