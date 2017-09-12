package windows

import (
	"goo/common"

	termbox "github.com/nsf/termbox-go"
)

type Window struct {
	Id         uint8
	Content    [][]byte
	Dimensions common.Dimensions
	Cursor     Cursor
}

func (window *Window) Insert(key termbox.Key, ch rune) {
	if ch == 0 && key == 0 {
		return
	}
	tempArr := window.Content[window.Cursor.Y]
	newLine := tempArr[:window.Cursor.X]
	newLine = append(newLine, byte(ch))
	newLine = append(newLine, tempArr[window.Cursor.X:]...)
	window.Content[window.Cursor.Y] = newLine
	window.MoveCursorRight()
}

func (window *Window) Draw() {
	for l, line := range window.Content {
		for c, char := range line {
			termbox.SetCell(c, l, rune(char), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
