package windows

import (
	"goo/common"
	"goo/window_events"
	"strconv"

	termbox "github.com/nsf/termbox-go"
)

type Window struct {
	Id                uint8
	Content           [][]byte
	DisableInsertMode bool
	EventHandlers     []window_events.WindowEvent
	EnableLineNum     bool
	CustomLoopFunc    func(*Window)
	Dimensions        common.Dimensions
	Position          common.Position
	OffsetH           int
	OffsetV           int
	Cursor            Cursor
}

func (window *Window) Insert(key termbox.Key, ch rune) {
	if ch == 0 && key == 0 {
		return
	}
	if key == termbox.KeyBackspace2 {
		window.backspace()
		return
	}

	if key == termbox.KeyEnter {
		window.enter()
		return
	}

	cursor := window.ContentCursor()
	newLine := common.SliceInsertIndex(window.Content[cursor.Y], byte(ch), cursor.X)

	window.Content[cursor.Y] = newLine
	window.MoveCursorRight()
}

func (window *Window) Draw() {
	position := window.Position
	verticalOffset := window.OffsetV
	horizontalOffset := window.OffsetH
	lineNumberMax := len(strconv.Itoa(len(window.Content)))
	size := window.Dimensions

	for l, line := range window.Content {
		window.drawBorders()
		window.drawLineNumbers()
		for c, char := range line {
			if c >= horizontalOffset && c <= size.Cols+horizontalOffset {
				if l >= verticalOffset && l <= size.Rows+verticalOffset {
					termbox.SetCell(c+position.X-horizontalOffset+lineNumberMax, l+position.Y-verticalOffset, rune(char), termbox.ColorDefault, termbox.ColorDefault)
				}
			}
		}
	}
}

func (window *Window) drawBorders() {
	position := window.Position
	for i := 0; i <= window.Dimensions.Cols; i++ {
		termbox.SetCell(i+position.X, position.Y-1, rune('-'), termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 0; i <= window.Dimensions.Rows-1; i++ {
		termbox.SetCell(position.X-1, i+position.Y, rune('|'), termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 0; i <= window.Dimensions.Rows-1; i++ {
		termbox.SetCell(position.X+window.Dimensions.Cols+1, i+position.Y, rune('|'), termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 0; i <= window.Dimensions.Cols; i++ {
		termbox.SetCell(i+position.X, position.Y+window.Dimensions.Rows+1, rune('-'), termbox.ColorDefault, termbox.ColorDefault)
	}

}

func (window *Window) drawLineNumbers() {
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
	return len(strconv.Itoa(len(window.Content)))
}
