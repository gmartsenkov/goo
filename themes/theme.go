package themes

import termbox "github.com/nsf/termbox-go"

type Theme struct {
	Background   termbox.Attribute
	TextColour   termbox.Attribute
	BorderColour termbox.Attribute
}
