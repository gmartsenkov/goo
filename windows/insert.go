package windows

import (
	"goo/common"

	termbox "github.com/nsf/termbox-go"
)

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

	cursor := window.ContentCursor()
	newLine := common.SliceInsertIndex(window.Content[cursor.Y], byte(ch), cursor.X)

	window.Content[cursor.Y] = newLine
	window.MoveCursorRight()
}
