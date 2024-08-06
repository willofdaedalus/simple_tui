package main

import (
	"fmt"

	"github.com/willofdaedalus/simple-tui/table"
)

func main() {
	t := table.NewTable([]string{"id", "item", "price", "quantity"}).
		AddRow([]string{"0", "apple", "0.50", "60"}).
		AddRow([]string{"0", "apple", "0.50", "60"})


	fmt.Print(t.DrawTable())
}
