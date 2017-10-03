package windows

import (
	"goo/common"
	"goo/window_events"
)

type Window struct {
	ID                    uint8
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
