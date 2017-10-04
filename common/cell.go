package common

import (
	termbox "github.com/nsf/termbox-go"
)

type Cell struct {
	Ch rune
	Bg termbox.Attribute
	Fg termbox.Attribute
}

type Cells [][]Cell
type CellsArray []Cell

func BytesToCells(b []byte) []Cell {
	tmp := []Cell{}
	for _, x := range b {
		tmp = append(tmp, Cell{Ch: rune(x)})
	}

	return tmp
}
func RunesToCellsWithStyle(style termbox.Attribute, b []rune) []Cell {
	tmp := []Cell{}
	for _, x := range b {
		tmp = append(tmp, Cell{Ch: x, Fg: style})
	}

	return tmp
}

func RunesToCells(b []rune) []Cell {
	return RunesToCellsWithStyle(0, b)
}

func CellsAsRuneArray(cells []Cell) []rune {
	runeArray := []rune{}
	for _, cell := range cells {
		runeArray = append(runeArray, cell.Ch)
	}

	return runeArray
}
