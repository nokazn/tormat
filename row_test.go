package tormat

import (
	"reflect"
	"testing"

	"github.com/nokazn/tormat/utils"
)

func TestRowGetMaxLength(t *testing.T) {
	type input struct {
		row Row
		min int
	}

	utils.RunTests(t, utils.TestRunner[input, int]{
		Name: "Row.getMaxLength",
		Cases: []utils.TestCase[input, int]{
			{
				Input:    input{Row{"", ""}, -1},
				Expected: 0,
			},
			{
				Input:    input{Row{"", ""}, 2},
				Expected: 2,
			},
			{
				Input:    input{Row{}, 0},
				Expected: 0,
			},
			{
				Input:    input{Row{}, 2},
				Expected: 2,
			},
			{
				Input:    input{Row{"favorite", "music", "avalanche"}, 2},
				Expected: 9,
			},
		},
		Runner: func(input input, expected int) (int, bool) {
			res := input.row.getMaxLength(input.min)
			ok := reflect.DeepEqual(res, expected)
			return res, ok
		},
	})
}

func TestRowNormalize(t *testing.T) {
	type input struct {
		length int
		row    Row
	}
	utils.RunTests(t, utils.TestRunner[input, Row]{
		Name: "Row.normalize",
		Cases: []utils.TestCase[input, Row]{
			{
				Input: input{
					length: -1,
					row:    Row{"1", "2", "3"},
				},
				Expected: Row{},
			},
			{
				Input: input{
					length: 0,
					row:    Row{"1", "2", "3"},
				},
				Expected: Row{},
			},
			{
				Input: input{
					length: 1,
					row:    Row{"1", "2", "3"},
				},
				Expected: Row{"1"},
			},
			{
				Input: input{
					length: 2,
					row:    Row{"1", "2", "3"},
				},
				Expected: Row{"1", "2"},
			},
			{
				Input: input{
					length: 3,
					row:    Row{"1", "2", "3"},
				},
				Expected: Row{"1", "2", "3"},
			},
			{
				Input: input{
					length: 4,
					row:    Row{"1", "2", "3"},
				},
				Expected: Row{"1", "2", "3", ""},
			},
			{
				Input: input{
					length: 10,
					row:    Row{"1", "2", "3"},
				},
				Expected: Row{"1", "2", "3", "", "", "", "", "", "", ""},
			},
		},
		Runner: func(input input, expected Row) (Row, bool) {
			res := input.row.normalize(input.length)
			ok := reflect.DeepEqual(res, expected)
			return res, ok
		},
	})
}
