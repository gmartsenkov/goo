package windows

import termbox "github.com/nsf/termbox-go"

func (window *Window) Draw() {
	position := window.Position
	verticalOffset := window.OffsetV
	horizontalOffset := window.OffsetH
	lineNumberMax := window.lineNumerLen()
	size := window.Dimensions

	if window.EnableSolidForeground {
		window.drawSolidForeground()
	}

	if window.EnableBorder {
		window.drawBorders()
	}
	if window.EnableLineNum {
		window.drawLineNumbers()
	}

	for l, line := range window.Content {
		for c, cell := range line {
			if c >= horizontalOffset && c <= size.Cols+horizontalOffset {
				if l >= verticalOffset && l <= size.Rows+verticalOffset {
					termbox.SetCell(c+position.X-horizontalOffset+lineNumberMax, l+position.Y-verticalOffset, cell.Ch, window.textStyle(), termbox.ColorDefault)
				}
			}
		}
	}
}

func (window *Window) drawSolidForeground() {
	pos := window.Position
	dimensions := window.Dimensions

	for y := pos.Y; y <= pos.Y+dimensions.Rows; y++ {
		for x := pos.X; x <= pos.X+dimensions.Cols; x++ {
			termbox.SetCell(x, y, rune(' '), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func (window *Window) textStyle() termbox.Attribute {
	if window.EnableBoldContent {
		return termbox.ColorDefault | termbox.AttrBold
	}
	return termbox.ColorDefault
}
