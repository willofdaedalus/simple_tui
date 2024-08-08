package table

import (
	"fmt"
	"strings"
)

// formatRowWrap recursively formats a row for rendering
func formatRowWrap(t *table, row []string) string {
	var b strings.Builder
	var remainder, printables []string

	for _, header := range row {
		// determine chop point
		chopPoint := t.width
		if len(header) > t.width {
			spaceIndex := strings.LastIndex(header[:t.width], " ")
			if spaceIndex != -1 {
				chopPoint = spaceIndex + 1 // include space in current line
			} else {
				chopPoint = t.width
			}
		}

		// ensure chopPoint does not exceed the header's length
		if chopPoint > len(header) {
			chopPoint = len(header)
		}

		// prepare current line and remainder
		currentPrint := header[:chopPoint]
		remainingPart := header[chopPoint:]

		// ensure the current line fits within the width and is properly padded
		if len(currentPrint) < t.width {
			currentPrint += strings.Repeat(" ", t.width-len(currentPrint))
		}
		printables = append(printables, currentPrint)

		// remainder should carry over
		if len(remainingPart) > 0 {
			remainder = append(remainder, remainingPart)
		} else {
			remainder = append(remainder, strings.Repeat(" ", t.width))
		}
	}

	// build the current line
	b.WriteByte('|')
	for _, s := range printables {
		b.WriteString(fmt.Sprintf(" %s |", s))
	}
	b.WriteByte('\n')

	// recursively handle remainder
	if hasNonEmptyEntries(remainder) {
		b.WriteString(formatRowWrap(t, remainder))
	}

	return b.String()
}

func hasNonEmptyEntries(row []string) bool {
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
