package events

import (
	"goo/editors"
	"goo/menu"

	termbox "github.com/nsf/termbox-go"
)

func EventLoop(menu *menu.Menu, editor *editors.Editor) {
	editor.Clear()

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

			for _, event := range CoreEvents {
				event.Process(ev.Key, editor)
			}
			for _, event := range VimEvents {
				event.Process(ev.Ch, editor)
			}

			menu.TriggerListener(editor, ev.Key, ev.Ch)
		}
	}

	editor.Draw()
	termbox.Flush()
}
