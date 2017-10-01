package windows

import (
	"bytes"
	"goo/common"
	"goo/window_events"
)

type Window struct {
	Id                    uint8
	Content               [][]byte
	DisableInsertMode     bool
	EnableSolidForeground bool
	EventHandlers         []window_events.WindowEvent
	EnableLineNum         bool
	EnableBorder          bool
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
		window.Content = append(window.Content, append([]byte{}, line...))
	}
}
