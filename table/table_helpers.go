package table

import (
	"fmt"
	"strings"
)

func (t *table) drawHeader(headers []string) string {
	var b strings.Builder
	longest := t.width

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
	longest := t.width
	totalLen := ((3 + longest) * len(row)) + 1

	for i, cell := range row {
		if i < colCount {
			str := cell
			if len(cell) > t.width && t.shorten {
				str = fmt.Sprintf("%s...", cell[:t.width-4])
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
