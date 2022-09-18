package utils

import (
	"reflect"
	"regexp"
	"runtime"
)

func GetFunctionName(f any) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	r := regexp.MustCompile(`.*\.(.+)$`)
	return r.ReplaceAllString(fullName, "$1")
}
