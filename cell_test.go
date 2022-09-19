package tormat

import (
	"testing"

	"github.com/nokazn/tormat/utils"
)

func TestCellPadEnd(t *testing.T) {
	type input struct {
		cell   Cell
		length uint
	}
	utils.RunTests(t, utils.TestRunner[input, Cell]{
		Name: "Cell.padEnd",
		Cases: []utils.TestCase[input, Cell]{
			{
				Input:    input{Cell(""), 0},
				Expected: Cell(""),
			},
			{
				Input:    input{Cell(" "), 0},
				Expected: Cell(" "),
			},
			{
				Input:    input{Cell(" foo  "), 0},
				Expected: Cell(" foo  "),
			},
			{
				Input:    input{Cell("foo"), 3},
				Expected: Cell("foo"),
			},
			{
				Input:    input{Cell("foo"), 5},
				Expected: Cell("foo  "),
			},
			{
				Input:    input{Cell("foo"), 10},
				Expected: Cell("foo       "),
			},
			{
				Input:    input{Cell(" foo "), 10},
				Expected: Cell(" foo      "),
			},
		},
		Runner: func(input input, expected Cell) (Cell, bool) {
			res := input.cell.padEnd(input.length)
			ok := res == expected
			return res, ok
		},
	})
}
