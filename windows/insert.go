package windows

import (
	"goo/common"

	termbox "github.com/nsf/termbox-go"
)

func (window *Window) Insert(key termbox.Key, ch rune) {

	if window.hitStandardKey(key, ch) {
		return
	}

	cursor := window.ContentCursor()
	newLine := common.SliceInsertIndex(window.Content[cursor.Y], byte(ch), cursor.X)

	window.Content[cursor.Y] = newLine
	window.ForceMoveCursorRight()
}

func (window *Window) hitStandardKey(key termbox.Key, ch rune) bool {
	if ch == 0 && key == 0 {
		return true
	}
	if key == termbox.KeyBackspace2 {
		window.backspace()
		return true
	}

	if key == termbox.KeyEnter {
		window.enter()
		return true
	}

	if key == termbox.KeyArrowUp {
		window.MoveCursorUp()
		return true
	}

	if key == termbox.KeyArrowDown {
		window.MoveCursorDown()
		return true
	}

	if key == termbox.KeyArrowLeft {
		window.MoveCursorLeft()
		return true
	}

	if key == termbox.KeyArrowRight {
		window.MoveCursorRight()
		return true
	}

	return false
}
