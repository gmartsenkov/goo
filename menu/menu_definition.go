package menu

var MENU = Menu{
	TriggerKey: 32,
	SubMenus: SubMenus{
		SubMenu{
			Title: "View",
			Key:   rune('v'),
			subMenus: SubMenus{
				SubMenu{
					Title: "Test",
					Key:   rune('t'),
				},
			},
		},
	},
}
