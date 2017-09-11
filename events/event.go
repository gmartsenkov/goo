package events

import (
	"goo/editors"

	"github.com/nsf/termbox-go"
)

type Event struct {
	Key termbox.Key
	fn  func(*editors.Editor)
}

func (event *Event) Process(key termbox.Key, editor *editors.Editor) {
	if event.Key == key {
		event.fn(editor)
	}
}

var Events = []Event{
	Event{
		Key: termbox.KeyEsc,
		fn:  func(w *editors.Editor) { w.Close() },
	},
}
