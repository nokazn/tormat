package tormat

import (
	"reflect"
	"testing"

	"github.com/nokazn/tormat/utils"
)

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
| 3       | 4       |`,
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
			ok := reflect.DeepEqual(res, expected)
			return res, ok
		},
	})
}
