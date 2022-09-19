package tormat

import (
	"regexp"
	"strings"

	"github.com/nokazn/tormat/utils"
)

type Column struct {
	header    Cell
	delimiter Cell
	body      []Cell
}

type ParsedTable struct {
	columns []Column
}

const (
	BAR            = "|"
	SPACE          = " "
	DELIMITER      = "-"
	NEWLINE        = "\n"
	NEWLINE_REGEXP = "\r?\n"
	MIN_DELIMITER  = 3
)

func parseRow(input string) Row {
	row := Input(input).toRow()
	l := len(row)
	if l == 0 {
		return Row{}
	}
	first := row[0]
	if l == 1 {
		return Row{first}
	}
	last := row[l-1]
	row = row[1 : l-1]
	if first != "" {
		row = utils.Unshift(row, first)
	}
	if last != "" {
		row = append(row, last)
	}
	return row.trim()
}

func parseTable(table string) (ParsedTable, error) {
	rows := regexp.MustCompile(NEWLINE_REGEXP).Split(strings.TrimSpace(table), -1)
	parsedTable := make(Table, len(rows))
	for y, row := range rows {
		parsedTable[y] = parseRow(row)
	}
	result, err := parsedTable.validate()
	if err != nil {
		return ParsedTable{}, err
	}
	body := result.body.rotate()
	return ParsedTable{
		columns: utils.Map(result.header, func(cell Cell, x uint) Column {
			return Column{
				header: cell,
				body:   body[x],
			}
		}),
	}, nil
}

func formatTable(table ParsedTable) ParsedTable {
	columns := utils.Map(table.columns, func(column Column, x uint) Column {
		size := append(Row{column.header}, column.body...).getMaxLength(MIN_DELIMITER)
		return Column{
			header:    column.header.padEnd(size),
			delimiter: Cell(strings.Repeat(DELIMITER, int(size))),
			body: utils.Map(column.body, func(cell Cell, y uint) Cell {
				return cell.padEnd(size)
			}),
		}
	})
	return ParsedTable{
		columns: columns,
	}
}

func stringifyTable(table ParsedTable) string {
	header := utils.Map(table.columns, func(column Column, x uint) Cell {
		return column.header
	})
	delimiter := utils.Map(table.columns, func(column Column, x uint) Cell {
		return column.delimiter
	})
	body := Table(utils.Map(table.columns, func(column Column, x uint) Row {
		return column.body
	})).rotate()
	rows := utils.Map(append(Table{header, delimiter}, body...), func(row Row, _ uint) string {
		return row.stringify()
	})
	return strings.Join(rows, NEWLINE)
}

/*
 * フォーマットする
 */
func Format(table string) string {
	parsedTable, err := parseTable(table)
	if err != nil {
		return table
	}
	return stringifyTable(formatTable(parsedTable))
}
