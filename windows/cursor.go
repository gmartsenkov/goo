package windows

import (
	"goo/common"
)

type Cursor struct {
	X int
	Y int
}

func (window *Window) ContentCursor() Cursor {
	lineNumLen := window.lineNumerLen()
	return Cursor{
		X: window.Cursor.X - window.Position.X + window.OffsetH - lineNumLen,
		Y: window.Cursor.Y - window.Position.Y + window.OffsetV,
	}
}

func (window *Window) MoveCursorDown() {
	cursor := window.ContentCursor()

	if cursor.Y >= len(window.Content)-1 {
		return
	}

	if cursor.Y >= window.Dimensions.Rows {
		window.OffsetV++
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
	cursor := window.ContentCursor()

	if cursor.Y <= 0 {
		return
	}

	if cursor.Y <= window.OffsetV {
		window.OffsetV--
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
	cursor := window.ContentCursor()

	if cursor.X <= 0 {
		return
	}

	if cursor.X <= window.OffsetH {
		window.OffsetH--
		return
	}
	window.Cursor.X--
}

func (window *Window) ForceMoveCursorRight() {
	window.Cursor.X++
}

func (window *Window) MoveCursorRight() {
	cursor := window.ContentCursor()

	lineLen := len(window.Content[cursor.Y]) - 1
	if cursor.X >= lineLen {
		return
	}

	if cursor.X >= window.Dimensions.Cols {
		window.OffsetH++
		return
	}
	window.Cursor.X++
}

func (window *Window) SetCursor(x, y int) {
	lineNumLen := window.lineNumerLen()

	window.Cursor.X = x + window.Position.X - window.OffsetH + lineNumLen
	window.Cursor.Y = y + window.Position.Y - window.OffsetV
}
