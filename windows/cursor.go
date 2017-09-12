package windows

type Cursor struct {
	X int
	Y int
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
