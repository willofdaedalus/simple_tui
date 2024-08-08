package main

import (
	"fmt"

	"github.com/willofdaedalus/simple-tui/table"
)

func main() {
	t2 := table.NewTable([]string{"member", "real name", "super power"}).
		Title("superheroes I guess").
		AddRow([]string{"spiderman", "peter parker"}).
		AddRow([]string{"", "", ""}).
		AddRow([]string{"hawkeye", "clint barton", "arrow powers"}).
		AddRow([]string{"ironman", "tony stark", "rich powers"})

	for i := 1; i <= 14; i += 1 {
		t2.SetWidth(i).
			Title(fmt.Sprintf("i: %d", i))
		fmt.Println(t2.DrawTable())
	}

}
