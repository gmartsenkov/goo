package main

import (
	"goo/common"
	"goo/editors"
	"goo/events"
	"goo/menu"
	"goo/windows"
	"io/ioutil"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	mainEditor := new(editors.Editor)
	menu := menu.MENU
	powerline(mainEditor)
	fileEditor(mainEditor)

	termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
	mainEditor.Draw()
	termbox.Flush()
	for {
		events.EventLoop(&menu, mainEditor)
	}
}

func fileEditor(mainEditor *editors.Editor) {
	w, h := termbox.Size()

	mainEditor.NewWindow(&windows.Window{
		EnableLineNum: true,
	})
	editor := mainEditor.CurrentWindow()
	editor.Dimensions = common.Dimensions{
		Cols: w,
		Rows: h - 3,
	}
	editor.Position = common.Position{
		X: 0,
		Y: 0,
	}
	editor.CustomLoopFunc = func(w *windows.Window) {
	}
	editor.SetCursor(1, 0)
	file, err := ioutil.ReadFile("main.go")
	if err != nil {
		panic("file not found")
	}
	editor.SplitAndSetContent(file)
}

func powerline(editor *editors.Editor) {
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
		w.Content = [][]byte{[]byte(editor.StateInWords())}
	}
}
