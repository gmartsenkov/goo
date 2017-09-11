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
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			for _, event := range events.CoreEvents {
				event.Process(ev.Key, mainEditor)
			}
			for _, event := range events.VimEvents {
				event.Process(ev.Ch, mainEditor)
			}
		}
	}
}
