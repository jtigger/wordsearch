package textui

import (
	"fmt"
	"os"
	"wordsearch/pkg/grid"
)

type UI interface {
	Render(grid.Grid)
}
var _ UI = (*ui)(nil)

type ui struct {
	in *os.File
	out *os.File
}

func (u ui) Render(g grid.Grid) {
	for _, row := range g.Letters {
		for _, cell := range row {
			fmt.Printf("%c  ", cell)
		}
		fmt.Printf("\n")
	}
}

func NewUI(in *os.File, out *os.File) UI {
   return &ui{in, out}
}