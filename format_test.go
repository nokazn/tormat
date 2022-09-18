package tormat

import (
	"reflect"
	"testing"

	"github.com/nokazn/tormat/utils"
)

func TestGetMaxLength(t *testing.T) {
	type Input struct {
		row Row
		min int
	}

	cases := []utils.TestCase[Input, int]{
		{
			Input:    Input{Row{"", ""}, -1},
			Expected: 0,
		},
		{
			Input:    Input{Row{"", ""}, 2},
			Expected: 2,
		},
		{
			Input:    Input{Row{"favorite", "music", "avalanche"}, 2},
			Expected: 9,
		},
	}

	utils.RunTests(t, utils.TestRunner[Input, int]{
		Fn:    getMaxLength,
		Cases: cases,
		Runner: func(input Input, expected int) (int, bool) {
			res := getMaxLength(input.row, input.min)
			ok := reflect.DeepEqual(res, expected)
			return res, ok
		}})
}
