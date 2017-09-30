package menu

import "errors"

type SubMenu struct {
	Title    string
	Key      rune
	subMenus SubMenus
	actions  Actions
}

type SubMenus []SubMenu

func (subMenus SubMenus) FindByKey(key rune) (SubMenu, error) {
	for _, subMenu := range subMenus {
		if subMenu.Key == key {
			return subMenu, nil
		}
	}
	return SubMenu{}, errors.New("SubMenuNotFound")
}
