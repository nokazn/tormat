package utils

import (
	"strconv"
	"testing"
)

type TestCase[I any, E any] struct {
	Input    I
	Expected E
}

type runner[I any, E any] func(input I, expected E) (E, bool)

type TestRunner[I any, E any] struct {
	Fn     any
	Cases  []TestCase[I, E]
	Runner runner[I, E]
}

func RunTests[I any, E any](t *testing.T, r TestRunner[I, E]) {
	funcName := GetFunctionName(r.Fn)
	for i, c := range r.Cases {
		t.Run(funcName+" "+strconv.Itoa(i), func(t *testing.T) {
			res, ok := r.Runner(c.Input, c.Expected)
			if ok {
				t.Logf("[input]: %v  [output]: %v", c.Input, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", c.Input, res, c.Expected)
			}
		})
	}
}
