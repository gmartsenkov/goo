package menu

import (
	"goo/common"
	"goo/editors"
	"goo/predefined_windows"

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
						content := common.Cells{}
						content = append(content, common.BytesToCells([]byte("")))
						content = append(content, common.BytesToCells([]byte(" Are you sure you want to quit?")))
						yN := append([]common.Cell{}, common.RunesToCells([]rune("            "))...)
						yN = append(yN, common.Cell{Ch: 'y', Fg: termbox.ColorGreen})
						yN = append(yN, common.Cell{Ch: '/'})
						yN = append(yN, common.Cell{Ch: 'n', Fg: termbox.ColorRed})
						yN = append(yN, common.RunesToCells([]rune("            "))...)

						w := predefined_windows.PopupCenterWindow(append(content, yN))

						drawBox := func() {
							e.Clear()
							e.DrawWindows()
							w.Draw(e.Theme)
							termbox.Flush()
						}

						drawBox()
						for {
							switch ev := termbox.PollEvent(); ev.Type {
							case termbox.EventKey:
								if ev.Ch == rune('y') {
									e.Close()
									return
								}
								if ev.Ch == rune('n') {
									e.Clear()
									e.DrawWindows()
									termbox.Flush()
									return
								}
							}
							drawBox()
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
