package events

import "goo/editors"
import "goo/common"

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
		Key: common.KeyJ,
		fn:  func(w *editors.Editor) { w.CurrentWindow().MoveCursorDown() },
	},
	VimEvent{
		Key: common.KeyK,
		fn:  func(w *editors.Editor) { w.CurrentWindow().MoveCursorUp() },
	},
	VimEvent{
		Key: common.KeyH,
		fn:  func(w *editors.Editor) { w.CurrentWindow().MoveCursorLeft() },
	},
	VimEvent{
		Key: common.KeyL,
		fn:  func(w *editors.Editor) { w.CurrentWindow().MoveCursorRight() },
	},
}
