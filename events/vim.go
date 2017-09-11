package events

import "goo/editors"

type VimEvent struct {
	Key rune
	fn  func(*editors.Editor)
}

func (event *VimEvent) Process(key rune, editor *editors.Editor) {
	if event.Key == key {
		event.fn(editor)
	}
}

var VimEvents = []VimEvent{
	VimEvent{
		Key: 106,
		fn:  func(w *editors.Editor) { w.CurrentWindow().MoveCursorDown() },
	},
	VimEvent{
		Key: 107,
		fn:  func(w *editors.Editor) { w.CurrentWindow().MoveCursorUp() },
	},
	VimEvent{
		Key: 104,
		fn:  func(w *editors.Editor) { w.CurrentWindow().MoveCursorLeft() },
	},
	VimEvent{
		Key: 108,
		fn:  func(w *editors.Editor) { w.CurrentWindow().MoveCursorRight() },
	},
}
