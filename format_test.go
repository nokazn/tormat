package tormat

import (
	"reflect"
	"testing"

	"github.com/nokazn/tormat/utils"
)

func TestGetMaxLength(t *testing.T) {
	type input struct {
		row Row
		min int
	}

	utils.RunTests(t, utils.TestRunner[input, int]{
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

func TestRotateMatrix(t *testing.T) {
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

func TestParseRow(t *testing.T) {
	utils.RunTests(t, utils.TestRunner[string, Row]{
		Fn: parseRow,
		Cases: []utils.TestCase[string, Row]{
			{
				Input:    "11|21|31|41",
				Expected: Row{"11", "21", "31", "41"},
			},
			{
				Input:    "|11|21|31|41|",
				Expected: Row{"11", "21", "31", "41"},
			},
			{
				Input:    "11|21|31|41|",
				Expected: Row{"11", "21", "31", "41"},
			},
			{
				Input:    " 11 |     21| 31   |41|              ",
				Expected: Row{"11", "21", "31", "41"},
			},
			{
				Input:    "11",
				Expected: Row{"11"},
			},
			{
				Input:    "11|",
				Expected: Row{"11"},
			},
			{
				Input:    "11| ",
				Expected: Row{"11"},
			},
			{
				Input:    "|",
				Expected: Row{},
			},
			{
				Input:    " |   ",
				Expected: Row{},
			},
			{
				Input:    "",
				Expected: Row{},
			},
		},
		Runner: func(input string, expected Row) (Row, bool) {
			res := parseRow(input)
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

func TestFormat(t *testing.T) {
	utils.RunTests(t, utils.TestRunner[string, string]{
		Fn: Format,
		Cases: []utils.TestCase[string, string]{
			{
				Input: `| header1 | header2 |
| ------- | --- |
| 1  | 2          |
| 3     | 4 |`,
				Expected: `| header1 | header2 |
| ------- | ------- |
| 1       | 2       |
| 3       | 4       |`,
			},
			{
				Input: `| header1 | header2 |
|------ | --- |
| 1  |          |
| 3     | 4 |`,
				Expected: `| header1 | header2 |
| ------- | ------- |
| 1       |         |
| 3       | 4       |
`,
			},
			{
				Input: `|             artist name|    url|
| --- | ------ |
|Tame Impala                  | https://tameimpala.com/        |
|     cocteau twins|  https://cocteautwins.com        |
|Fishmans|                               http://www.fishmans.jp/                             |`,
				Expected: `| artist name   | url                      |
| ------------- | ------------------------ |
| Tame Impala   | https://tameimpala.com/  |
| cocteau twins | https://cocteautwins.com |
| Fishmans      | http://www.fishmans.jp/  |`,
			},
			{
				Input: `


header1 | header2 | header3 | | |
| --- | --- |--- | --- | ---|
1-1 | 1-2 | 1-3  | 1-4 |         1 -  5
| 2-1 | 2 -  2 |          |2-4|2-5
3-1 | 3-2 |  | | 3-5|
|
|| 4- 2 |  | |


			`,
				Expected: `| header1 | header2 | header3 |     |        |
| ------- | ------- | ------- | --- | ------ |
| 1-1     | 1-2     | 1-3     | 1-4 | 1 -  5 |
| 2-1     | 2 -  2  |         | 2-4 | 2-5    |
| 3-1     | 3-2     |         |     | 3-5    |
|         |         |         |     |        |
|         | 4- 2    |         |     |        |`,
			},
			{
				Input: `||||||
|---|---|---|---|---|
||||||
|||||
||||||
||||||
||||||
||||||
||||||
|
||||||
||||||`,
				Expected: `|     |     |     |     |     |
| --- | --- | --- | --- | --- |
|     |     |     |     |     |
|     |     |     |     |     |
|     |     |     |     |     |
|     |     |     |     |     |
|     |     |     |     |     |
|     |     |     |     |     |
|     |     |     |     |     |
|     |     |     |     |     |
|     |     |     |     |     |
|     |     |     |     |     |`,
			},
		},
		Runner: func(input string, expected string) (string, bool) {
			res := Format(input)
			ok := reflect.DeepEqual(input, expected)
			return res, ok
		},
	})
}
