package windows

import (
	"bytes"
	"goo/common"
	"goo/window_events"
)

type Window struct {
	Id                    uint8
	Content               common.Cells
	DisableInsertMode     bool
	EnableSolidForeground bool
	EventHandlers         []window_events.WindowEvent
	EnableLineNum         bool
	EnableBorder          bool
	EnableBoldContent     bool
	CustomLoopFunc        func(*Window)
	Dimensions            common.Dimensions
	Position              common.Position
	OffsetH               int
	OffsetV               int
	Cursor                Cursor
}

func (window *Window) SplitAndSetContent(content []byte) {
	tmp := bytes.Split(content, []byte("\n"))

	for _, line := range tmp {
		window.Content = append(window.Content, common.BytesToCells(line))
	}
}

func (window *Window) ContentAsRuneArray() [][]rune {
	tmp := [][]rune{}
	for _, line := range window.Content {
		tmp = append(tmp, common.CellsAsRuneArray(line))
	}

	return tmp
}

func (window *Window) AppendLineRuneArray(line []rune) {
	window.Content = append(window.Content, common.RunesToCells(line))
}
