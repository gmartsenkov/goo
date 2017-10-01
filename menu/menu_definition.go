package menu

import (
	"goo/common"
	"goo/editors"
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
						if w.EnableLineNum {
							w.EnableLineNum = false
						} else {
							w.EnableLineNum = true
						}
					},
				},
			},
		},
		SubMenu{
			Title: "Editor",
			Key:   rune('e'),
			subMenus: SubMenus{
				SubMenu{
					Title: "States",
					Key:   rune('b'),
					actions: Actions{
						Action{
							Title: "-> Insert State",
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
