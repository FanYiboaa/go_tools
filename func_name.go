package tools

import (
	"reflect"
	"runtime"
)

// GetFuncName 获取函数名
func GetFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
