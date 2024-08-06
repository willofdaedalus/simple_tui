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
	Width   int
	Shorten bool
	Rows    [][]string
}

func NewTable(row []string) *table {
	t := &table{
		Rows:  make([][]string, 0),
		Width: 10,
	}
	t.Rows = append(t.Rows, row)
	return t
}

func (t *table) AddRow(row []string) *table {
	t.Rows = append(t.Rows, row)
	return t
}

func (t *table) AddRows(rows [][]string) *table {
	for _, row := range rows {
		t.Rows = append(t.Rows, row)
	}
	return t
}

func (t *table) SetWidth(w int) *table {
	t.Width = w
	return t
}

func (t *table) Truncate(b bool) *table {
	t.Shorten = b
	return t
}

func DrawHeader(s string, truncate bool, c int, decorations int) string {
	var b strings.Builder
	var l, r byte

	str := s
	if truncate {
		if c == 0 {
			c = 15
		}
		if len(s) > c {
			str = fmt.Sprintf("%s...", s[:c])
		}
	}

	strlen := len(str) + 4

	if decorations&LEFT_BORDER != 0 {
		l = '|'
		strlen += 1
	}

	if decorations&RIGHT_BORDER != 0 {
		r = '|'
		strlen += 1
	}

	b.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", strlen)))
	b.WriteString(fmt.Sprintf("%c  %s  %c", l, str, r))
	b.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", strlen)))

	return b.String()
}

func (t *table) drawHeaders(headers []string) string {
	var b strings.Builder
	longest := t.Width

	// calculate total length of the header line
	totalLen := ((3 + longest) * len(headers)) + 1

	// build the top border
	b.WriteString(strings.Repeat("-", totalLen))
	b.WriteString("\n")

	// build the header row
	for _, header := range headers {
		padding := strings.Repeat(" ", longest-len(header))
		b.WriteString(fmt.Sprintf("| %s%s ", header, padding))
	}
	b.WriteString("|\n")

	// build the bottom border
	b.WriteString(strings.Repeat("-", totalLen))
	b.WriteString("\n")

	return b.String()
}

func (t *table) drawRow(row []string, colCount int, end bool) string {
	var builder strings.Builder
	longest := t.Width
	totalLen := ((3 + longest) * len(row)) + 1

	for i, cell := range row {
		if i < colCount {
			str := cell
			if len(cell) > t.Width {
				str = fmt.Sprintf("%s...", cell[:t.Width-4])
			}

			padding := strings.Repeat(" ", longest-len(cell))
			builder.WriteString(fmt.Sprintf("| %s%s ", str, padding))
		}
	}
	builder.WriteString("|\n")
	if end {
		builder.WriteString(strings.Repeat("-", totalLen))
		builder.WriteString("\n")
	}

	return builder.String()
}

func (t *table) DrawTable() string {
	var builder strings.Builder
	var columnCount int

	for i, row := range t.Rows {
		if i == 0 {
			builder.WriteString(t.drawHeaders(row))
			columnCount = len(row)
		} else if i == len(t.Rows)-1 {
			builder.WriteString(t.drawRow(row, columnCount, true))
		} else {
			builder.WriteString(t.drawRow(row, columnCount, false))
		}
	}

	return builder.String()
}
