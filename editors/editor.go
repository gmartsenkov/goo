package editors

import (
	"goo/common"
	"goo/windows"
	"os"

	termbox "github.com/nsf/termbox-go"
)

type Editor struct {
	windows       []*windows.Window
	dimensions    common.Dimensions
	currentWindow uint8
}

func (editor *Editor) Draw() {
	for _, buffer := range editor.windows {
		buffer.Draw()
	}
	cursor := editor.CurrentWindow().Cursor
	termbox.SetCursor(cursor.X, cursor.Y)
}

func (editor *Editor) CurrentWindow() *windows.Window {
	for _, window := range editor.windows {
		if editor.currentWindow == window.Id {
			return window
		}
	}

	return &windows.Window{}
}

func (editor *Editor) NewWindow() {
	window := &windows.Window{
		Id: uint8(len(editor.windows)),
	}
	editor.windows = append(editor.windows, window)
}

func (editor *Editor) Close() {
	termbox.Close()
	os.Exit(0)
}
