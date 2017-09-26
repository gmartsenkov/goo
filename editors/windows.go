package editors

import (
	"goo/windows"
	"os"

	termbox "github.com/nsf/termbox-go"
)

func (editor *Editor) CurrentWindow() *windows.Window {
	for _, window := range editor.Windows {
		if editor.currentWindow == window.Id {
			return window
		}
	}

	return &windows.Window{}
}

func (editor *Editor) NewWindow(window *windows.Window) {
	id := uint8(len(editor.Windows) + 1)
	window.Id = id
	editor.currentWindow = id
	editor.Windows = append(editor.Windows, window)
}

func (editor *Editor) Close() {
	termbox.Close()
	os.Exit(0)
}
