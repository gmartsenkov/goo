package editors

import (
	"goo/common"
	"goo/windows"
	"os"

	termbox "github.com/nsf/termbox-go"
)

type Editor struct {
	buffers    []*windows.Window
	dimensions common.Dimensions
}

func (editor *Editor) NewWindow() {
}

func (editor *Editor) Close() {
	termbox.Close()
	os.Exit(0)
}
