package tormat

import "strings"

type Input string

func (input Input) toRow() Row {
	trimmed := strings.Trim(string(input), " ")
	if trimmed == "" {
		return Row{}
	}
	cells := strings.Split(trimmed, BAR)
	row := make(Row, len(cells))
	for x, v := range cells {
		row[x] = Cell(v)
	}
	return row
}
