package tormat

import "strings"

type Cell string

func (cell Cell) padEnd(length int) Cell {
	shortage := length - len(cell)
	if shortage < 1 {
		return cell
	}
	return cell + Cell(strings.Repeat(SPACE, shortage))
}
