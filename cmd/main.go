package main

import (
	"os"
	"strings"
	"wordsearch/pkg/textui"
	"wordsearch/pkg/grid"
)

func main() {
	//grid, _ := grid.Generate(strings.NewReader("hard,large,garden,march,shark,yard,car,star,smart,farm"))
	grid, _ := grid.Generate(strings.NewReader("july,january,august,skeleton,philosophical,niagrafalls,ninja"))
	ui := textui.NewUI(os.Stdin, os.Stdout)
	ui.Render(grid)
}