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
	mainEditor.NewWindow()
	powerline := mainEditor.CurrentWindow()
	powerline.Dimensions = common.Dimensions{
		Cols: 100,
		Rows: 1,
	}
	powerline.Position = common.Position{
		X: 0,
		Y: 20,
	}
	powerline.CustomLoopFunc = func(w *windows.Window) {
		w.Content = [][]byte{[]byte(mainEditor.StateInWords())}
	}
	mainEditor.NewWindow()
	editor := mainEditor.CurrentWindow()
	editor.Dimensions = common.Dimensions{
		Cols: 50,
		Rows: 10,
	}
	editor.Position = common.Position{
		X: 2,
		Y: 2,
	}
	editor.CustomLoopFunc = func(w *windows.Window) {
	}
	editor.Cursor.X = 3
	editor.Cursor.Y = 2
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
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte("clablax3"),
		[]byte(""),
	}

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
