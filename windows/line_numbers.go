package windows

import (
	"goo/common"
	"goo/themes"
	"strconv"

	termbox "github.com/nsf/termbox-go"
)

func (window *Window) drawLineNumbers(theme themes.Theme) {
	position := window.Position
	lineLenNum := window.lineNumerLen()
	num := make([]byte, lineLenNum)

	for i := 1; i <= common.MinInt(len(window.Content), window.Dimensions.Rows)+1; i++ {
		num = []byte(strconv.Itoa(i + window.OffsetV))

		for x, digit := range num {
			termbox.SetCell(position.X+x, i+position.Y-1, rune(digit), termbox.ColorBlack, termbox.ColorWhite)
		}

		for x := len(num); x < lineLenNum; x++ {
			termbox.SetCell(position.X+x, i+position.Y-1, rune(' '), termbox.ColorBlack, termbox.ColorWhite)
		}
	}
}

func (window *Window) lineNumerLen() int {
	if window.EnableLineNum {
		return len(strconv.Itoa(len(window.Content)))
	}
	return 0
}
