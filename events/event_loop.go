package events

import (
	"goo/editors"
	"goo/menu"
	"goo/windows"

	termbox "github.com/nsf/termbox-go"
)

func EventLoop(menu *menu.Menu, editor *editors.Editor) {
	termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)

	menuWindow := &windows.Window{}
	menuWindow.EnableBorder = true
	menuWindow.EnableSolidForeground = true

	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		if editor.State == editors.StateInsert {
			if ev.Key == termbox.KeyEsc {
				editor.NormalState()
				ev.Key = 0
			}
			editor.CurrentWindow().Insert(ev.Key, ev.Ch)
		}
		if editor.State == editors.StateNormal {
			w, h := termbox.Size()
			menuWindow.Dimensions.Cols = w

			if menu.IsTriggered() {
				menu.Process(menuWindow, editor, ev.Ch)
				menuWindow.Position.Y = h - menuWindow.Dimensions.Rows
			} else {
				for _, event := range CoreEvents {
					event.Process(ev.Key, editor)
				}
				for _, event := range VimEvents {
					event.Process(ev.Ch, editor)
				}
			}

			menu.TriggerListener(menuWindow, ev.Key)
			menuWindow.Dimensions.Rows = len(menuWindow.Content)
			menuWindow.Position.Y = h - menuWindow.Dimensions.Rows - 3
		}
	}

	editor.Draw()
	if menuWindow.Content != nil {
		menuWindow.Draw()
	}
	termbox.Flush()
}
