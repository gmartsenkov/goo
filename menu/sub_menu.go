package menu

import (
	"errors"
	"goo/common"
)

type SubMenu struct {
	Title    string
	Key      rune
	subMenus SubMenus
	actions  Actions
}

type element struct {
	key    rune
	title  string
	symbol rune
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

func (subMenus SubMenus) ContentForWindow(submenu SubMenu, width int) [][]common.Cell {
	content := [][]common.Cell{[]common.Cell{}}
	elements := []element{}
	row := 0

	for _, subMenu := range subMenus {
		elements = append(elements, element{
			key:    subMenu.Key,
			title:  subMenu.Title,
			symbol: '\u2295',
		})
	}

	for _, action := range submenu.actions {
		elements = append(elements, element{
			key:    action.Key,
			title:  action.Title,
			symbol: '\u25ce',
		})
	}

	for _, element := range elements {
		x := []common.Cell{}
		x = append(x, common.Cell{Ch: rune(' ')})
		x = append(x, common.Cell{Ch: rune(' ')})
		for _, c := range []rune(element.title) {
			x = append(x, common.Cell{Ch: c})
		}
		x = append(x, common.Cell{Ch: rune('\u2794')})
		x = append(x, common.Cell{Ch: rune(' ')})
		x = append(x, common.Cell{Ch: element.key})

		if len(content[row])+len(x) >= width {
			content = append(content, x)
			row++
		} else {
			content[row] = append(content[row], x...)
		}
	}

	return content
}
