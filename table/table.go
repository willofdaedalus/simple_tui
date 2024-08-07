package table

import (
	"fmt"
	"strings"
)

const (
	NO_BORDER    = 0
	LEFT_BORDER  = 1 << 1
	RIGHT_BORDER = 1 << 2
)

type table struct {
	title      string
	width      int
	truncate   bool
	longestRow int
	rows       [][]string
}

func NewTable(row []string) *table {
	t := &table{
		rows:  make([][]string, 0),
		width: 10,
	}
	t.rows = append(t.rows, row)
	return t
}

func (t *table) AddRow(row []string) *table {
	t.rows = append(t.rows, row)
	return t
}

func (t *table) AddRows(rows [][]string) *table {
	for _, row := range rows {
		t.rows = append(t.rows, row)
	}
	return t
}

func (t *table) SetWidth(w int) *table {
	if w > 0 {
		t.width = w
	}
	return t
}

func (t *table) Title(s string) *table {
	t.title = s
	return t
}

func (t *table) noTruncate(b bool) *table {
	t.truncate = b
	return t
}

func (t *table) DrawTable() {
	var builder strings.Builder
	var columnCount int

	if len(t.title) != 0 {
		fmt.Println(t.title)
	}

	for i, row := range t.rows {
		if i == 0 {
			builder.WriteString(t.drawHeader(row))
			columnCount = len(row)
		} else if i == len(t.rows)-1 {
			builder.WriteString(t.drawRow(row, columnCount, true))
		} else {
			builder.WriteString(t.drawRow(row, columnCount, false))
		}
	}

	fmt.Println(builder.String())
}
