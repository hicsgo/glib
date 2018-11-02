package glib

import (
	"errors"
	"github.com/lunny/log"
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
			log.Error("invoke func panic error:", err)
		}
	}()
	fnSource()
}

