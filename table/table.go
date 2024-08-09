package table

import (
	"fmt"
	"strings"
)

type table struct {
	title      string     // title for the table
	width      int        // width of each individual cell on a row
	rows       [][]string // all rows in the table
	dynamicRows bool
}

// NewTable creates a new table that can be populated
func NewTable(row []string) *table {
	t := &table{
		rows:  make([][]string, 0),
		width: 10,
	}
	t.rows = append(t.rows, row)
	return t
}

// AddRow adds a single row to the table
func (t *table) AddRow(row []string) *table {
	t.rows = append(t.rows, row)
	return t
}

// AddRows adds rows to the table. Currently untested
func (t *table) AddRows(rows [][]string) *table {
	for _, row := range rows {
		t.rows = append(t.rows, row)
	}
	return t
}

func (t *table) Dynamic(b bool) *table {
	t.dynamicRows = b
	return t
}

// SetWidth sets the width of each cell in each row
func (t *table) SetWidth(w int) *table {
	if w > 0 {
		t.width = w
	}
	return t
}

// Title sets a title that printed above the table when DrawTable() is called
func (t *table) Title(s string) *table {
	t.title = s
	return t
}

// DrawTable formats and returns the table in ASCII format ready for printing
func (t *table) DrawTable() string {
	var builder strings.Builder
	var rowRep string
	var columnCount int

	if len(t.title) != 0 {
		fmt.Println(t.title)
	}

	for i, row := range t.rows {
		// first slice is always considered the header
		if i == 0 {
			rowRep = t.drawHeader(row)

			builder.WriteString(drawBorderLine(rowRep))
			builder.WriteString(rowRep)
			builder.WriteString(drawBorderLine(rowRep))

			columnCount = len(row)
			continue
		}

		builder.WriteString(t.drawRow(row, columnCount))
	}

	// draw the final line at the bottom
	builder.WriteString(drawBorderLine(rowRep))

	return builder.String()
}
