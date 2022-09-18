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
