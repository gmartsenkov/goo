package predefined_windows

import (
	"goo/common"
	"goo/editors"
	"goo/windows"
	"io/ioutil"

	termbox "github.com/nsf/termbox-go"
)

func TextEditor(editor *editors.Editor) {
	w, h := termbox.Size()

	editor.NewWindow(&windows.Window{
		EnableLineNum: true,
	})

	textEditor := editor.CurrentWindow()
	textEditor.Dimensions = common.Dimensions{
		Cols: w,
		Rows: h - 3,
	}
	textEditor.Position = common.Position{
		X: 0,
		Y: 0,
	}
	textEditor.CustomLoopFunc = func(w *windows.Window) {
	}

	file, err := ioutil.ReadFile("main.go")
	if err != nil {
		panic("file not found")
	}
	textEditor.SplitAndSetContent(file)
	textEditor.SetCursor(1, 0)
}
