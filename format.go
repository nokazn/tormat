package tormat

type Cell string
type Row []Cell
type Table []Row

func getMaxLength(row Row, min int) int {
	max := min
	for _, v := range row {
		l := len(v)
		if max < l {
			max = l
		}
	}
	return max
}

func rotateTable(table Table) Table {
	if len(table) == 0 {
		return []Row{}
	}
	firstRow := table[0]
	if len(firstRow) == 0 {
		return []Row{}
	}
	rotatedTable := make(Table, len(firstRow))
	for x := range firstRow {
		rotatedTable[x] = make(Row, len(table))
		for y := range table {
			rotatedTable[x][y] = table[y][x]
		}
	}
	return rotatedTable
}
