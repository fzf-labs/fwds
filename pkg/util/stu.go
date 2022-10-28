package util

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var Stu = NewStu()

type stu struct {
}

func NewStu() *stu {
	return &stu{}
}

func (s *stu) StructToJsonTagValue(u interface{}) string {
	var str string
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for i := 0; i < t.NumField(); i++ {
		str += t.Field(i).Tag.Get("json") + "=" + fmt.Sprintf("%v", v.Field(i).Interface()) + "&"
	}
	st := strings.Trim(str, "&")
	query, err := url.ParseQuery(st)
	if err != nil {
		return ""
	}
	return query.Encode()
}

// StructToSliceAndMap 结构体转换为 切片key,切片值,map
func (s *stu) StructToSliceAndMap(data interface{}) (sliceKey *[]string, sliceVale *[]string, Map *map[string]string) {
	kvs := make(map[string]string)
	var keys []string
	var vals []string
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		field := t.Field(i)
		tag := string(field.Tag)
		reg, err := regexp.Compile(`json:"(.*?)"`)
		if err != nil {
			continue
		}
		rs := reg.FindStringSubmatch(tag)
		if len(rs) != 2 {
			continue
		}
		tagName := rs[1]
		typeString := field.Type.String()
		va := v.Field(i).Interface()
		var value string
		if typeString == "string" {
			value = va.(string)
		} else if typeString == "int64" {
			value = strconv.FormatInt(va.(int64), 10)
		} else if typeString == "int" || typeString == "int32" {
			value = strconv.Itoa(va.(int))
		} else if typeString == "float64" {
			value = strconv.FormatFloat(va.(float64), 'f', 30, 32)
		} else if typeString == "bool" {
			value = strconv.FormatBool(va.(bool))
		} else if typeString == "uint64" {
			value = strconv.FormatUint(va.(uint64), 10)
		} else if typeString == "unit" || typeString == "uint32" {
			value = strconv.Itoa(va.(int))
		}
		if value == "" {
			continue
		}
		keys = append(keys, tagName)
		vals = append(vals, value)
		kvs[tagName] = value
	}
	return &keys, &vals, &kvs
}
