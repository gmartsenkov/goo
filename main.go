package main

import (
	"fmt"
	"goo/editors"
	"goo/events"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	h, w := termbox.Size()
	mainEditor := new(editors.Editor)

	fmt.Println(h, w)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(1, 2, rune('b'), termbox.ColorDefault, termbox.ColorWhite)
	termbox.SetCell(2, 3, rune('a'), termbox.ColorDefault, termbox.ColorWhite)
	termbox.Flush()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			for _, event := range events.Events {
				event.Process(ev.Key, mainEditor)
			}
		}
	}
}
