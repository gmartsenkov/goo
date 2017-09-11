package events

import (
	"goo/editors"

	termbox "github.com/nsf/termbox-go"
)

func EventLoop(editor *editors.Editor) {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		for _, event := range CoreEvents {
			event.Process(ev.Key, editor)
		}
		if editor.State == editors.StateNormal {
			for _, event := range VimEvents {
				event.Process(ev.Ch, editor)
			}
		}
	}
}
