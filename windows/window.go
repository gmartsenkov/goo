package windows

import (
	"goo/common"

	termbox "github.com/nsf/termbox-go"
)

type Window struct {
	Id         uint8
	Content    []byte
	Dimensions common.Dimensions
	Cursor     Cursor
}

func (window *Window) MoveCursorDown() {
	window.Cursor.Y++
}

func (window *Window) MoveCursorUp() {
	window.Cursor.Y--
}

func (window *Window) MoveCursorLeft() {
	window.Cursor.X--
}

func (window *Window) MoveCursorRight() {
	window.Cursor.X++
}

func (window *Window) Draw() {
	window.Content = []byte("blabla\n bla bla \n")
	line := 0
	row := 0
	for _, char := range window.Content {
		if char == byte('\n') {
			line = line + 1
			row = 0
			termbox.SetCell(row, line, rune(char), termbox.ColorDefault, termbox.ColorDefault)
		} else {
			termbox.SetCell(row, line, rune(char), termbox.ColorDefault, termbox.ColorDefault)
			row++
		}
	}
}
