package editors

import (
	"goo/common"
	"goo/themes"
	"goo/windows"

	termbox "github.com/nsf/termbox-go"
)

type Editor struct {
	Theme         themes.Theme
	Windows       []*windows.Window
	dimensions    common.Dimensions
	State         uint8
	currentWindow uint8
}

func (editor *Editor) Clear() {
	termbox.Clear(editor.Theme.TextColour, editor.Theme.Background)
}

func (editor *Editor) Draw() {
	editor.DrawWindows()
	cursor := editor.CurrentWindow().Cursor
	termbox.SetCursor(cursor.X, cursor.Y)
}

func (editor *Editor) DrawWindows() {
	for _, buffer := range editor.Windows {
		buffer.CustomLoopFunc(buffer)
		buffer.Draw(editor.Theme)
	}
}

func (editor *Editor) Size() common.Dimensions {
	w, h := termbox.Size()
	return common.Dimensions{
		Cols: w,
		Rows: h,
	}
}
