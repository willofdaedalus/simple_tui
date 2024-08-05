package main

import (
	"fmt"

	"github.com/willofdaedalus/simple-tui/table"
)

func main() {
	fmt.Print(table.DrawHeader("hello world", true, 0, table.RIGHT_BORDER | table.LEFT_BORDER))
	fmt.Print(table.DrawHeader("this is a test", true, 0, table.LEFT_BORDER))
	fmt.Print(table.DrawHeader("what's this?", true, 0, table.RIGHT_BORDER))
}
