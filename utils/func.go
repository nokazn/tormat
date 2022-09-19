package utils

import (
	"errors"
	"reflect"
	"regexp"
	"runtime"
)

func GetFunctionName(f any) (string, error) {
	if f == nil {
		return "", errors.New("関数がありません")
	}
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	r := regexp.MustCompile(`.*\.(.+)$`)
	return r.ReplaceAllString(fullName, "$1"), nil
}
