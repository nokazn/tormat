package tormat

import (
	"reflect"
	"testing"

	"github.com/nokazn/tormat/utils"
)

func TestTableRotate(t *testing.T) {
	utils.RunTests(t, utils.TestRunner[Table, Table]{
		Name: "Table.rotate",
		Cases: []utils.TestCase[Table, Table]{
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
		},
		Runner: func(input Table, expected Table) (Table, bool) {
			res := input.rotate()
			ok := reflect.DeepEqual(res, expected)
			return res, ok
		},
	})
}
