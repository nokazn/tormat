package tormat

import (
	"math"
	"strings"

	"github.com/nokazn/tormat/utils"
)

type Row []Cell

/*
 * 各セルの両端のスペースを詰める
 */
func (row Row) trim() Row {
	for x, v := range row {
		row[x] = Cell(strings.TrimSpace(string(v)))
	}
	return row
}

/*
 * 列の長さを揃える
 */
func (row Row) normalize(length int) Row {
	shortage := int(math.Max(float64(length), 0)) - len(row)
	if shortage >= 0 {
		return append(row, make(Row, shortage)...)
	}
	return row[:len(row)+shortage]
}

/*
 * 要素の最大の長さを求める
 */
func (row Row) getMaxLength(min int) int {
	max := min
	for _, v := range row {
		l := len(v)
		if max < l {
			max = l
		}
	}
	return max
}

/*
 * 行を文字列に戻す
 */
func (row Row) stringify() string {
	list := append([]string{BAR}, utils.Map(row, func(cell Cell, x int) string {
		return string(cell) + SPACE + BAR
	})...)
	return strings.Join(list, SPACE)
}
