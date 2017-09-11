package windows

import (
	"goo/common"

	termbox "github.com/nsf/termbox-go"
)

type Window struct {
	content    []byte
	dimensions common.Dimensions
}

func (window *Window) Close() {
	termbox.Close()
}
