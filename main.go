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

	termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
	for {
		mainEditor.Draw()
		termbox.Flush()
		events.EventLoop(mainEditor)
	}
}
