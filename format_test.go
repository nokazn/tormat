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
		},
	})
}

func TestRotateMatrix(t *testing.T) {
	cases := []utils.TestCase[Table, Table]{
		{
			Input: Table{
				Row{"11", "21", "31", "41"},
			},
			Expected: Table{
				Row{"11"},
				Row{"21"},
				Row{"31"},
				Row{"41"},
			},
		},
		{
			Input: Table{
				Row{"11", "21", "31"},
				Row{"12", "22", "32"},
				Row{"13", "23", "33"},
			},
			Expected: Table{
				Row{"11", "12", "13"},
				Row{"21", "22", "23"},
				Row{"31", "32", "33"},
			},
		},
		{
			Input: Table{
				Row{"11", "21", "31", "41"},
				Row{"12", "22", "32", "42"},
				Row{"13", "23", "33", "43"},
			},
			Expected: Table{
				Row{"11", "12", "13"},
				Row{"21", "22", "23"},
				Row{"31", "32", "33"},
				Row{"41", "42", "43"},
			},
		},
		{
			Input:    Table{},
			Expected: Table{},
		},
		{
			Input: Table{
				Row{},
			},
			Expected: Table{},
		},
	}

	utils.RunTests(t, utils.TestRunner[Table, Table]{
		Fn:    rotateTable,
		Cases: cases,
		Runner: func(input Table, expected Table) (Table, bool) {
			res := rotateTable(input)
			ok := reflect.DeepEqual(res, expected)
			return res, ok
		},
	})
}
