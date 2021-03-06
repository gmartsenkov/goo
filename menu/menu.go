package menu

import (
	"errors"
	"goo/editors"
	"goo/predefined_windows"
	"goo/windows"

	termbox "github.com/nsf/termbox-go"
)

type Menu struct {
	TriggerKey termbox.Key
	triggered  bool
	SubMenus   SubMenus
	Query      []rune
}

func (menu *Menu) TriggerListener(editor *editors.Editor, key termbox.Key, ch rune) {
	w, h := termbox.Size()
	window := predefined_windows.Menu()
	window.Dimensions.Cols = w
	drawMenu := func() {
		editor.Clear()
		editor.DrawWindows()
		window.Draw(editor.Theme)
		termbox.Flush()
	}

	if key == menu.TriggerKey {
		menu.triggered = true
		window.Content = menu.SubMenus.ContentForWindow(SubMenu{}, window.Dimensions.Cols)
		window.Dimensions.Rows = len(window.Content)
		window.Position.Y = h - window.Dimensions.Rows - 3
		drawMenu()
	}

	for menu.IsTriggered() {

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			menu.Process(&window, editor, ev.Ch)
		}

		window.Dimensions.Rows = len(window.Content)
		window.Position.Y = h - window.Dimensions.Rows - 3
		drawMenu()
	}

	editor.Clear()
	editor.Draw()
	termbox.Flush()
}

func (menu *Menu) IsTriggered() bool {
	return menu.triggered == true
}

func (menu *Menu) Process(window *windows.Window, editor *editors.Editor, key rune) {
	subMenu, err := menu.FindSubMenu(key)
	if err == nil {
		window.Content = subMenu.subMenus.ContentForWindow(subMenu, window.Dimensions.Cols)
		menu.Query = append(menu.Query, key)
		return
	}

	subMenu, err = menu.FindLastSubMenu()
	if err == nil {
		action, err := subMenu.actions.FindByKey(key)
		if err == nil {
			action.Fn(editor)
			menu.Query = []rune{}
			menu.triggered = false
			return
		}
	}

	menu.Query = []rune{}
	menu.triggered = false
}

func (menu *Menu) FindSubMenu(k rune) (SubMenu, error) {
	subMenu := menu.SubMenus
	for _, key := range menu.Query {
		tempSubMenu, err := subMenu.FindByKey(rune(key))
		if err == nil {
			subMenu = tempSubMenu.subMenus
		}
	}
	s, err := subMenu.FindByKey(k)
	if err == nil {
		return s, nil
	}
	return SubMenu{}, errors.New("SubMenuNotFound")
}

func (menu *Menu) FindLastSubMenu() (SubMenu, error) {
	subMenus := menu.SubMenus
	lastSubmenu := SubMenu{}

	for _, key := range menu.Query {
		tempSubMenus, err := subMenus.FindByKey(rune(key))
		if err == nil {
			subMenus = tempSubMenus.subMenus
			lastSubmenu = tempSubMenus
		}
	}

	if lastSubmenu.Title != "" {
		return lastSubmenu, nil
	}

	return SubMenu{}, errors.New("SubMenuNotFound")
}
