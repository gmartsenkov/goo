package windows

import (
	"goo/common"
	"goo/themes"

	termbox "github.com/nsf/termbox-go"
)

func (window *Window) Draw(theme themes.Theme) {
	position := window.Position
	verticalOffset := window.OffsetV
	horizontalOffset := window.OffsetH
	lineNumberMax := window.lineNumerLen()
	size := window.Dimensions

	if window.EnableSolidForeground {
		window.drawSolidForeground(theme)
	}

	if window.EnableBorder {
		window.drawBorders(theme)
	}
	if window.EnableLineNum {
		window.drawLineNumbers(theme)
	}

	for l, line := range window.Content {
		for c, cell := range line {
			if c >= horizontalOffset && c <= size.Cols+horizontalOffset {
				if l >= verticalOffset && l <= size.Rows+verticalOffset {
					termbox.SetCell(c+position.X-horizontalOffset+lineNumberMax, l+position.Y-verticalOffset, cell.Ch, window.textStyle(theme, cell), theme.Background)
				}
			}
		}
	}
}

func (window *Window) drawSolidForeground(theme themes.Theme) {
	pos := window.Position
	dimensions := window.Dimensions

	for y := pos.Y; y <= pos.Y+dimensions.Rows; y++ {
		for x := pos.X; x <= pos.X+dimensions.Cols; x++ {
			termbox.SetCell(x, y, rune(' '), theme.TextColour, theme.Background)
		}
	}
}

func (window *Window) textStyle(theme themes.Theme, cell common.Cell) termbox.Attribute {
	style := theme.TextColour

	if cell.Fg != 0 {
		style = cell.Fg
	}

	if window.EnableBoldContent {
		style = style | termbox.AttrBold
	}

	return style
}
