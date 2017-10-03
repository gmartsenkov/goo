package menu

import (
	"goo/common"
	"goo/editors"
	"goo/windows"

	termbox "github.com/nsf/termbox-go"
)

var MENU = Menu{
	TriggerKey: common.Space,
	SubMenus: SubMenus{
		SubMenu{
			Title: "Windows",
			Key:   rune('w'),
			actions: Actions{
				Action{
					Title: "Lines",
					Key:   rune('l'),
					Fn: func(e *editors.Editor) {
						w := e.CurrentWindow()
						w.EnableLineNum = !w.EnableLineNum
					},
				},
			},
		},
		SubMenu{
			Title: "Editor",
			Key:   rune('e'),
			actions: Actions{
				Action{
					Title: "Quit",
					Key:   rune('q'),
					Fn: func(e *editors.Editor) {
						w := windows.Window{}
						w.Content = append(w.Content, common.BytesToCells([]byte("Are you sure you want to quit?")))
						w.Content = append(w.Content, common.BytesToCells([]byte("            y/n               ")))
						w.Position.X = 20
						w.Position.Y = 10
						w.Dimensions.Cols = 30
						w.Dimensions.Rows = 10
						w.EnableBorder = true
						w.EnableSolidForeground = true
						w.EnableBoldContent = true

						e.DrawWindows()
						w.Draw()
						termbox.Flush()

						for {
							switch ev := termbox.PollEvent(); ev.Type {
							case termbox.EventKey:
								if ev.Ch == rune('y') {
									e.Close()
									return
								}
								if ev.Ch == rune('n') {
									termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
									e.DrawWindows()
									termbox.Flush()
									return
								}
							}
							e.DrawWindows()
							w.Draw()
							termbox.Flush()
						}
					},
				},
			},
			subMenus: SubMenus{
				SubMenu{
					Title: "States",
					Key:   rune('b'),
					actions: Actions{
						Action{
							Title: "Insert State",
							Key:   rune('m'),
							Fn: func(e *editors.Editor) {
								e.InsertState()
							},
						},
					},
				},
			},
		},
	},
}
