package windows

import termbox "github.com/nsf/termbox-go"

func (window *Window) drawBorders() {
	position := window.Position
	for i := 0; i <= window.Dimensions.Cols; i++ {
		termbox.SetCell(i+position.X, position.Y-1, rune('-'), termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 0; i <= window.Dimensions.Rows; i++ {
		termbox.SetCell(position.X-1, i+position.Y, rune('|'), termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 0; i <= window.Dimensions.Rows; i++ {
		termbox.SetCell(position.X+window.Dimensions.Cols+1, i+position.Y, rune('|'), termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 0; i <= window.Dimensions.Cols; i++ {
		termbox.SetCell(i+position.X, position.Y+window.Dimensions.Rows+1, rune('-'), termbox.ColorDefault, termbox.ColorDefault)
	}
}
