package tormat

import (
	"strings"

	"github.com/nokazn/tormat/utils"
)

type RawRow string
type Input string

func (input RawRow) toRow() Row {
	trimmed := strings.Trim(string(input), SPACE)
	if trimmed == "" {
		return Row{}
	}
	return utils.Map(
		strings.Split(trimmed, BAR),
		func(val string, _ uint) Cell { return Cell(val) },
	)
}
