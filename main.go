package main

import (
	"goo/editors"
	"goo/events"
	"goo/menu"
	"goo/predefined_windows"
	"goo/themes"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output256)
	mainEditor := new(editors.Editor)
	mainEditor.Theme = themes.DEFAULT

	menu := menu.MENU

	predefined_windows.Powerline(mainEditor)
	predefined_windows.TextEditor(mainEditor)

	mainEditor.Clear()
	mainEditor.Draw()
	termbox.Flush()
	for {
		events.EventLoop(&menu, mainEditor)
	}
}
