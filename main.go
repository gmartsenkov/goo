package main

import (
	"goo/editors"
	"goo/events"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	mainEditor := new(editors.Editor)
	mainEditor.NewWindow()
	mainEditor.CurrentWindow().Content = [][]byte{
		[]byte("blabla"),
		[]byte(" bla bla "),
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
