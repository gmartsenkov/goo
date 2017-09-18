package windows

import (
	"goo/common"
)

type Cursor struct {
	X int
	Y int
}

func (window *Window) MoveCursorDown() {
	cursor := window.Cursor

	if cursor.Y >= len(window.Content)-1 {
		return
	}

	nextLineLen := len(window.Content[cursor.Y+1])
	if cursor.X >= nextLineLen {
		window.SetCursor(common.MaxInt(nextLineLen-1, 0), cursor.Y+1)
		return
	}

	window.Cursor.Y++
}

func (window *Window) MoveCursorUp() {
	cursor := window.Cursor
	if cursor.Y <= 0 {
		return
	}

	previousLineLen := len(window.Content[cursor.Y-1])
	if cursor.X >= previousLineLen {
		window.SetCursor(common.MaxInt(previousLineLen-1, 0), cursor.Y-1)
		return
	}

	window.Cursor.Y--
}

func (window *Window) MoveCursorLeft() {
	if window.Cursor.X <= 0 {
		return
	}
	window.Cursor.X--
}

func (window *Window) MoveCursorRight() {
	cursor := window.Cursor
	if cursor.X >= len(window.Content[cursor.Y])-1 {
		return
	}
	window.Cursor.X++
}

func (window *Window) SetCursor(x, y int) {
	window.Cursor.X = x
	window.Cursor.Y = y
}
