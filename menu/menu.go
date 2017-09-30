package menu

import (
	"errors"
)

type Menu struct {
	TriggerKey rune // 32
	triggered  bool
	SubMenus   SubMenus
	Query      []byte
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
