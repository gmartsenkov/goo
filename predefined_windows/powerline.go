package predefined_windows

import (
	"goo/common"
	"goo/editors"
	"goo/windows"

	termbox "github.com/nsf/termbox-go"
)

func Powerline(editor *editors.Editor) {
	w, h := termbox.Size()
	editor.NewWindow(&windows.Window{
		EnableBorder: true,
	})

	powerline := editor.CurrentWindow()
	powerline.Dimensions = common.Dimensions{
		Cols: w,
		Rows: 1,
	}
	powerline.Position = common.Position{
		X: 0,
		Y: h - 1,
	}
	powerline.CustomLoopFunc = func(w *windows.Window) {
		w.Content = append(common.Cells{}, common.BytesToCells([]byte(editor.StateInWords())))
	}
}
