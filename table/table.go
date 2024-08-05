package table

import (
	"fmt"
	"strings"
)

const (
	NO_BORDER = 0
	LEFT_BORDER  = 1 << 1
	RIGHT_BORDER = 1 << 2
)

type table struct {
	Headers []string
	Rows    []string
}

func (t *table) AddRows(rows []string) {
}

func NewTable(headers []string) *table {
	return &table{
		Headers: headers,
	}
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

	if decorations & LEFT_BORDER != 0 {
		l = '|'
		strlen += 1
	}

	if decorations & RIGHT_BORDER != 0 {
		r = '|'
		strlen += 1
	}

	b.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", strlen)))
	b.WriteString(fmt.Sprintf("%c  %s  %c", l, str, r))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", strlen)))

	return b.String()
}

func (t *table) DrawTable() string {
	var builder strings.Builder

	return builder.String()
}
