package windows

import (
	"goo/common"
	"goo/window_events"

	termbox "github.com/nsf/termbox-go"
)

type Window struct {
	Id                uint8
	Content           [][]byte
	DisableInsertMode bool
	EventHandlers     []window_events.WindowEvent
	Dimensions        common.Dimensions
	Cursor            Cursor
}

func (window *Window) Insert(key termbox.Key, ch rune) {
	if ch == 0 && key == 0 {
		return
	}
	if key == termbox.KeyBackspace2 {
		window.backspace()
		return
	}

	if key == termbox.KeyEnter {
		window.enter()
		return
	}

	cursor := window.Cursor

	newLine := common.SliceInsertIndex(window.Content[cursor.Y], byte(ch), cursor.X)

	window.Content[cursor.Y] = newLine
	window.MoveCursorRight()
}

func (window *Window) Draw() {
	for l, line := range window.Content {
		for c, char := range line {
			termbox.SetCell(c, l, rune(char), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
