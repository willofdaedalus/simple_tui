package main

import (
	"github.com/willofdaedalus/simple-tui/table"
)

func main() {
	t := table.NewTable([]string{"id", "item", "price", "quantity"}).
		SetWidth(15).
		Truncate(true).
		AddRow([]string{"0", "apple", "0.50", "60"}).
		AddRow([]string{"1", "banana", "0.25", "76"}).
		AddRow([]string{"2", "orange", "0.35", "156"})

	t2 := table.NewTable([]string{"member", "real name", "super power"}).
	Title("superheroes I guess").
	SetWidth(15).
	AddRow([]string{"spiderman", "peter parker", "spider powers"}).
	AddRow([]string{"the hulk", "bruce banner", "anger powers"}).
	AddRow([]string{"hawkeye", "clint barton", "arrow powers"}).
	AddRow([]string{"ironman", "tony stark", "rich powers"})

	t.DrawTable()
	t2.DrawTable()

}
