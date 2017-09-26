package main

import (
	"goo/common"
	"goo/editors"
	"goo/events"
	"goo/windows"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	mainEditor := new(editors.Editor)
	powerline(mainEditor)
	fileEditor(mainEditor)

	termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
	mainEditor.Draw()
	termbox.Flush()
	for {
		events.EventLoop(mainEditor)
		termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
		mainEditor.Draw()
		termbox.Flush()
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
		Rows: h - 1,
	}
	editor.Position = common.Position{
		X: 0,
		Y: 0,
	}
	editor.CustomLoopFunc = func(w *windows.Window) {
	}
	editor.SetCursor(0, 0)
	editor.Content = [][]byte{
		[]byte("blabla"),
		[]byte(" blabla "),
		[]byte(" blablax3"),
		[]byte(" blablax3"),
		[]byte(" blablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
	}

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
