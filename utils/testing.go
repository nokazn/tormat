package utils

import (
	"strconv"
	"testing"
)

type TestCase[I any, E any] struct {
	Input    I
	Expected E
}

type TestRunner[I any, E any] struct {
	Fn     any
	Name   string
	Cases  []TestCase[I, E]
	Runner func(input I, expected E) (E, bool)
}

func RunTests[I any, E any](t *testing.T, r TestRunner[I, E]) {
	funcName, err := GetFunctionName(r.Fn)
	if err != nil {
		if r.Name != "" {
			funcName = r.Name
		} else {
			funcName = ""
		}
	}
	for i, c := range r.Cases {
		t.Run(funcName+" "+strconv.Itoa(i), func(t *testing.T) {
			res, ok := r.Runner(c.Input, c.Expected)
			if ok {
				t.Logf("[input]: %v\n[output]: %v", c.Input, res)
			} else {
				t.Errorf("[input]: %v\n[output]: %v\n[expected]: %v", c.Input, res, c.Expected)
			}
		})
	}
}
