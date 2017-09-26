package windows

import termbox "github.com/nsf/termbox-go"

func (window *Window) Draw() {
	position := window.Position
	verticalOffset := window.OffsetV
	horizontalOffset := window.OffsetH
	lineNumberMax := window.lineNumerLen()
	size := window.Dimensions

	if window.EnableBorder {
		window.drawBorders()
	}
	if window.EnableLineNum {
		window.drawLineNumbers()
	}

	for l, line := range window.Content {
		for c, char := range line {
			if c >= horizontalOffset && c <= size.Cols+horizontalOffset {
				if l >= verticalOffset && l <= size.Rows+verticalOffset {
					termbox.SetCell(c+position.X-horizontalOffset+lineNumberMax, l+position.Y-verticalOffset, rune(char), termbox.ColorDefault, termbox.ColorDefault)
				}
			}
		}
	}
}
