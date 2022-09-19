package tormat

import (
	"errors"
)

type Table []Row

type ValidationResult struct {
	header Row
	body   Table
}

/*
 * 行と列を入れ替える
 */
func (table Table) rotate() Table {
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

/*
 * 各行の列数を揃える
 */
func (table Table) normalize(length uint) Table {
	for y, row := range table {
		table[y] = row.normalize(length)
	}
	return table
}

/*
 * テーブルが正しい形式か検証する
 */
func (table Table) validate() (ValidationResult, error) {
	l := len(table)
	if l == 0 {
		return ValidationResult{}, errors.New("ヘッダが存在しません")
	}
	if l == 1 {
		return ValidationResult{}, errors.New("デリミタが存在しません")
	}
	header := table[0]
	delimiter := table[1]
	columnSize := len(header)
	if columnSize == 0 || columnSize != len(delimiter) {
		return ValidationResult{}, errors.New("テーブルの列数が不正です")
	}
	return ValidationResult{
		header: header,
		body:   table[2:].normalize(uint(columnSize)),
	}, nil
}
