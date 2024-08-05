package table

import (
	"fmt"
	"strings"
)

type table struct {
	Headers []string
	Rows []string
}

func (t *table) AddRows(rows []string) {
}

func NewTable(headers []string) *table {
	return &table{
		Headers: headers,
	}
}

func DrawHeader(s string, truncate bool, c int) string {
	var b strings.Builder

	str := s
	if truncate {
		if c == 0 {
			c = 15
		}
		str = fmt.Sprintf("%s...", s[:c])
	}

	l := len(str) + 6

	b.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", l)))
	b.WriteString(fmt.Sprintf("|  %s  |\n", str))
	b.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", l)))

	return b.String()
}

func (t *table) DrawTable() string {
	var builder strings.Builder



	return builder.String()
}
