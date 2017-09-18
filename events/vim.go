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
		fn: func(w *editors.Editor) {
			w.CurrentWindow().MoveCursorDown()
		},
	},
	VimEvent{
		Key: common.KeyShiftA,
		fn: func(w *editors.Editor) {
			window := w.CurrentWindow()
			cursor := window.Cursor

			lineCharCount := len(window.Content[cursor.Y])
			window.Content[cursor.Y] = append(window.Content[cursor.Y], byte(' '))

			window.SetCursor(lineCharCount, cursor.Y)
			w.InsertState()
		},
	},
	VimEvent{
		Key: common.KeyShiftI,
		fn: func(w *editors.Editor) {
			window := w.CurrentWindow()
			cursor := window.Cursor

			window.SetCursor(0, cursor.Y)
			w.InsertState()
		},
	},
	VimEvent{
		Key: common.KeyA,
		fn: func(w *editors.Editor) {
			window := w.CurrentWindow()
			cursor := window.Cursor
			if len(window.Content[cursor.Y])-1 == cursor.X {
				window.Content[cursor.Y] = append(window.Content[cursor.Y], byte(' '))
			}

			window.SetCursor(window.Cursor.X+1, window.Cursor.Y)
			w.InsertState()
		},
	},
	VimEvent{
		Key: common.KeyI,
		fn:  func(w *editors.Editor) { w.InsertState() },
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
