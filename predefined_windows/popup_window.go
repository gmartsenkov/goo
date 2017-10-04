package predefined_windows

import (
	"goo/common"
	"goo/windows"

	termbox "github.com/nsf/termbox-go"
)

func PopupCenterWindow(content common.Cells) *windows.Window {
	w, h := termbox.Size()

	window := &windows.Window{}
	window.Content = content

	window.Dimensions.Cols = maxCols(content)
	window.Dimensions.Rows = len(content)
	window.Position.X = (w / 2) - window.Dimensions.Cols/2
	window.Position.Y = (h / 2) - window.Dimensions.Rows/2

	window.EnableBorder = true
	window.EnableSolidForeground = true
	window.EnableBoldContent = true

	return window
}

func maxCols(cells common.Cells) int {
	x := 0
	for _, cell := range cells {
		cap := len(cell)
		if cap > x {
			x = cap
		}
	}

	return x
}
