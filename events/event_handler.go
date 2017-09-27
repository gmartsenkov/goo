package events

import (
	"errors"
	"goo/editors"
)

type Menu struct {
	TriggerKey rune // 32
	triggered  bool
	SubMenus   SubMenus
	Query      []byte
}

type SubMenu struct {
	Title    string
	Key      rune
	subMenus SubMenus
	actions  Actions
}

type Action struct {
	Title string
	Key   rune
	Fn    func(*editors.Editor)
}

type SubMenus []SubMenu
type Actions []Action

func (subMenus SubMenus) FindByKey(key rune) (SubMenu, error) {
	for _, subMenu := range subMenus {
		if subMenu.Key == key {
			return subMenu, nil
		}
	}
	return SubMenu{}, errors.New("Crap")
}

var X = Menu{
	TriggerKey: 32,
	SubMenus: SubMenus{
		SubMenu{
			Title: "View",
			Key:   rune('v'),
		},
	},
}
