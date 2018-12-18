package glib

import (
	"errors"
	"fmt"
	"strings"
	"encoding/json"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 捕获函数执行时的异常(外部函数可能有其他业务操作)
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Capture2(fnSource func(), fnError func(interface{})) (er error) {
	defer func() {
		if err := recover(); err != nil {
			fnError(err)
			if err, ok := err.(error); ok {
				er = err
			} else {
				er = errors.New("try func call errors")
			}
		}
	}()

	fnSource()
	return er
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 捕获函数执行时的异常
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Capture(fnSource func(...interface{})) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Sprintf("invoke func panic error:%v", err)
		}
	}()
	fnSource()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 捕获函数执行时的异常
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Capture3(fnSource func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Sprintf("invoke func panic error:%v", err)
		}
	}()
	fnSource()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用指定的字符串链接字符串切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringSliceToString(stringSlice []string, args ...string) string {
	result := ""

	if len(stringSlice) == 0 {
		return result
	}

	joinString := ","
	if len(args) == 1 {
		joinString = args[0]
	}

	if len(stringSlice) == 1 {
		result = strings.Join(stringSlice, "")
	} else {
		for _, v := range stringSlice {
			if len(result) == 0 {
				result = result + v
			} else {
				result = result + joinString + v
			}
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Json字符串转换成对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FromJson(jsonString string, object interface{}) error {
	bytesData := []byte(jsonString)
	return json.Unmarshal(bytesData, object)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转换成Json字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ToJson(object interface{}) (string, error) {
	v, err := json.Marshal(object)
	if err != nil {
		return "", err
	}

	return string(v), nil
}
