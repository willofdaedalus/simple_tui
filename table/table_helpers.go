package table

import (
	"fmt"
	"strings"
)

func (t *table) drawHeader(headers []string) string {
	var b strings.Builder
	// calculate total length of the header line
	offset := 2   // offset accounts for the spaces before and after each entry
	ellipsis := 3 // number of dots to append if shorten is true
	totalLen := ((offset + t.width + ellipsis) * len(headers)) + 1

	// build the top border
	b.WriteString(strings.Repeat("-", totalLen + 3))
	b.WriteString("\n")

	// build the header row
	for _, header := range headers {
		str := header
		strLen := len(header)
		pad := ""

		if strLen < t.width {
			pad = strings.Repeat("*", t.width - strLen)
		}

		b.WriteString(fmt.Sprintf("| %s%s ", str, pad))
	}
	b.WriteString("|\n")

	// build the bottom border
	b.WriteString(strings.Repeat("-", totalLen+3))
	b.WriteString("\n")

	return b.String()
}

func truncateText(s string, w int) (string, int) {
	var str string

	if len(str) > w {
		str = fmt.Sprintf("%s... ", s[:w])
	} else if (len(str) == w) || (len(str) < w) {
		str = fmt.Sprintf("%s%s ", s, strings.Repeat("*", w + 2))
	}

	return str, len(str)
}

func (t *table) drawRow(row []string, colCount int, end bool) string {
	var builder strings.Builder
	// calculate total length of the header line
	offset := 2   // offset accounts for the spaces before and after each entry
	ellipsis := 3 // number of dots to append if shorten is true
	totalLen := ((offset + t.width + ellipsis) * len(row)) + 1

	for i, cell := range row {
		if i < colCount {
			str := cell
			if len(cell) > t.width && t.truncate {
				str = fmt.Sprintf("%s...", cell[:t.width])
			}

			padding := "*"
			if len(cell) < t.width {
				// spread := t.width - len(dynamicStr)
				padding = strings.Repeat("*", t.width+ellipsis)
			}
			builder.WriteString(fmt.Sprintf("| %s%s", str, padding))
		}
	}
	builder.WriteString("|\n")
	if end {
		builder.WriteString(strings.Repeat("-", totalLen))
		builder.WriteString("\n")
	}

	return builder.String()
}
