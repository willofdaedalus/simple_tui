package table

import (
	"fmt"
	"strings"
)

// func (t *table) drawHeader(headers []string) string {
// 	var b strings.Builder
// 	var cutOff []string
// 	// calculate total length of the header line
// 	offset := 2   // offset accounts for the spaces before and after each entry
// 	ellipsis := 3 // number of dots to append if shorten is true
// 	totalLen := ((offset + t.width + ellipsis) * len(headers)) + 1
//
// 	// build the top border
// 	b.WriteString(strings.Repeat("-", totalLen + 3))
// 	b.WriteString("\n")
//
// 	// build the header row
// 	for _, header := range headers {
// 		str := header
// 		strLen := len(header)
// 		pad := ""
//
// 		if strLen < t.width {
// 			pad = strings.Repeat("*", t.width - strLen)
// 		} else {
// 			str = header[:t.width]
//
// 		}
//
// 		b.WriteString(fmt.Sprintf("| %s%s ", str, pad))
// 	}
// 	b.WriteString("|\n")
//
// 	// build the bottom border
// 	b.WriteString(strings.Repeat("-", totalLen+3))
// 	b.WriteString("\n")
//
// 	return b.String()
// }

// func (t *table) drawHeader(headers []string) string {
// 	var b strings.Builder
// 	var remainder, printables []string
//
// 	for _, header := range headers {
// 		curStr := header
//
// 		if len(curStr) > t.width {
// 			printables = append(printables, curStr[:t.width])
// 			remainder = append(remainder, curStr[t.width:])
// 		} else if len(curStr) <= t.width {
// 			printables = append(printables, curStr)
// 			remainder = append(remainder, strings.Repeat(" ", t.width))
// 		// } else if len(curStr) == 0 {
// 		// 	// should I append the empty strings here?
// 		// 	printables = append(printables, strings.Repeat(" ", t.width))
// 		}
// 	}
//
// 	for _, s := range printables {
// 		b.WriteString(fmt.Sprintf("| %s ", s))
// 	}
//
// 	b.WriteByte('\n')
//
// 	if len(remainder) > 0 {
// 		t.drawHeader(remainder)
// 	}
//
// 	return b.String()
// }

func (t *table) drawHeader(headers []string) string {
	var b strings.Builder
	var remainder, printables []string

	// process each header to split into current print and remainder
	for _, header := range headers {
		if len(header) > t.width {
			printables = append(printables, header[:t.width])
			remainder = append(remainder, header[t.width:])
		} else {
			header = fmt.Sprintf("%s%s", header, strings.Repeat(" ", t.width - len(header)))
			printables = append(printables, header)
			remainder = append(remainder, strings.Repeat(" ", t.width - len(header)))
		}
	}

	// build the current line
	b.WriteByte('|')
	for _, s := range printables {
		b.WriteString(fmt.Sprintf(" %s |", s))
	}
	b.WriteByte('\n')

	// if there's more to print, handle the remainder recursively
	if hasNonEmptyStrings(remainder) {
		b.WriteString(t.drawHeader(remainder))
	}

	return b.String()
}

// helper function to check if there are any non-empty strings in the remainder
func hasNonEmptyStrings(headers []string) bool {
	for _, header := range headers {
		if len(strings.TrimSpace(header)) > 0 {
			return true
		}
	}
	return false
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
