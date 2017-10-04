package editors

import (
	"goo/common"
	"goo/windows"

	termbox "github.com/nsf/termbox-go"
)

type Editor struct {
	Windows       []*windows.Window
	dimensions    common.Dimensions
	State         uint8
	currentWindow uint8
}

func (editor *Editor) Clear() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
}

func (editor *Editor) Draw() {
	editor.DrawWindows()
	cursor := editor.CurrentWindow().Cursor
	termbox.SetCursor(cursor.X, cursor.Y)
}

func (editor *Editor) DrawWindows() {
	for _, buffer := range editor.Windows {
		buffer.CustomLoopFunc(buffer)
		buffer.Draw()
	}
}
