package windows

import (
	"bytes"
	"goo/common"
	"goo/window_events"
)

type Window struct {
	Id                uint8
	Content           [][]byte
	DisableInsertMode bool
	EventHandlers     []window_events.WindowEvent
	EnableLineNum     bool
	EnableBorder      bool
	CustomLoopFunc    func(*Window)
	Dimensions        common.Dimensions
	Position          common.Position
	OffsetH           int
	OffsetV           int
	Cursor            Cursor
}

func (window *Window) SplitAndSetContent(content []byte) {
	window.Content = bytes.Split(content, []byte("\n"))
}
