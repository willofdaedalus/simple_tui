package main

import (
	"fmt"

	"github.com/willofdaedalus/simple-tui/table"
)

func main() {
	t2 := table.NewTable([]string{"member", "real name li 12234567", "super power"}).
		Title("8").
		SetWidth(8).
		AddRow([]string{"spiderman", "peter parker"}).
		AddRow([]string{"hawkeye", "clint barton", "arrow powers"}).
		AddRow([]string{"ironman", "tony stark", "rich powers"})

	fmt.Println(t2.DrawTable())

}
