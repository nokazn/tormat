package tormat

import "strings"

type Cell string

func (cell Cell) padEnd(length uint) Cell {
	shortage := int(length) - len(cell)
	if shortage < 1 {
		return cell
	}
	return cell + Cell(strings.Repeat(SPACE, shortage))
}
