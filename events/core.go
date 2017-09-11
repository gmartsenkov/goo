package events

import (
	"goo/editors"

	termbox "github.com/nsf/termbox-go"
)

type CoreEvent struct {
	Key termbox.Key
	fn  func(*editors.Editor)
}

func (event *CoreEvent) Process(key termbox.Key, editor *editors.Editor) {
	if event.Key == key {
		event.fn(editor)
	}
}

var CoreEvents = []CoreEvent{
	CoreEvent{
		Key: termbox.KeyEsc,
		fn:  func(w *editors.Editor) { w.Close() },
	},
	CoreEvent{
		Key: termbox.KeyCtrlC,
		fn:  func(w *editors.Editor) { w.Close() },
	},
}
