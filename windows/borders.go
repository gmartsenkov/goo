package windows

import (
	"goo/themes"

	termbox "github.com/nsf/termbox-go"
)

func (window *Window) drawBorders(theme themes.Theme) {
	position := window.Position

	for i := 0; i <= window.Dimensions.Cols; i++ {
		termbox.SetCell(i+position.X, position.Y-1, rune('\u2500'), theme.BorderColour, theme.Background)
	}
	for i := 0; i <= window.Dimensions.Rows; i++ {
		termbox.SetCell(position.X-1, i+position.Y, rune('\u2502'), theme.BorderColour, theme.Background)
	}

	for i := 0; i <= window.Dimensions.Rows; i++ {
		termbox.SetCell(position.X+window.Dimensions.Cols+1, i+position.Y, rune('\u2502'), theme.BorderColour, theme.Background)
	}
	for i := 0; i <= window.Dimensions.Cols; i++ {
		termbox.SetCell(i+position.X, position.Y+window.Dimensions.Rows+1, rune('\u2500'), theme.BorderColour, theme.Background)
	}

	termbox.SetCell(position.X+window.Dimensions.Cols+1, position.Y-1, rune('\u2510'), theme.BorderColour, theme.Background)
	termbox.SetCell(position.X+window.Dimensions.Cols+1, position.Y+window.Dimensions.Rows+1, rune('\u2518'), theme.BorderColour, theme.Background)
	termbox.SetCell(position.X-1, position.Y+window.Dimensions.Rows+1, rune('\u2514'), theme.BorderColour, theme.Background)
	termbox.SetCell(position.X-1, position.Y-1, rune('\u250c'), theme.BorderColour, theme.Background)
}
