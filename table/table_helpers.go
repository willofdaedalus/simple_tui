package table

import (
	"fmt"
	"strings"
)

// formatRowWrap recursively formats a row for rendering
func formatRowWrap(t *table, row []string) string {
	var b strings.Builder
	// remainder gets passed to the next recursive run of formatRowWrap
	// printables contains the strings that will be added to the strings builder
	var remainder, printables []string

	for _, header := range row {
		if len(header) > t.width {
			printables = append(printables, header[:t.width])
			remainder = append(remainder, header[t.width:])
		} else {
			paddedHeader := fmt.Sprintf("%s%s", header, strings.Repeat(" ", t.width - len(header)))
			printables = append(printables, paddedHeader)
			remainder = append(remainder, strings.Repeat(" ", t.width))
		}
	}

	// Build the current line
	b.WriteByte('|') // first | in the row
	for _, s := range printables {
		b.WriteString(fmt.Sprintf(" %s |", s))
	}
	b.WriteByte('\n') // move cursor to next line

	// recursively handle remainder
	if hasEmptyEntries(remainder) {
		b.WriteString(formatRowWrap(t, remainder))
	}

	return b.String()
}

// helper function to check if there are any non-empty strings in the remainder
func hasEmptyEntries(row []string) bool {
	for _, header := range row {
		if len(strings.TrimSpace(header)) > 0 {
			return true
		}
	}
	return false
}

// function for drawing borders
func drawBorderLine(s string) string {
	nl := strings.Index(s, "\n")
	return strings.Repeat("-", nl) + "\n"
}

// function specific for drawing row
func (t *table) drawHeader(row []string) string {
	return formatRowWrap(t, row)
}

// function specific for drawing rows
func (t *table) drawRow(row []string, colCount int) string {
	var builder strings.Builder
	var printable []string = make([]string, colCount)

	// maybe this is not the right way?
	copy(printable, row)
	builder.WriteString(formatRowWrap(t, printable))

	return builder.String()
}
