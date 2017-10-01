package menu

import "errors"
import "fmt"

type SubMenu struct {
	Title    string
	Key      rune
	subMenus SubMenus
	actions  Actions
}

type element struct {
	key   rune
	title string
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

func (subMenus SubMenus) ContentForWindow(submenu SubMenu, width int) [][]byte {
	content := [][]byte{[]byte{}}
	elements := []element{}
	row := 0

	for _, subMenu := range subMenus {
		elements = append(elements, element{key: subMenu.Key, title: subMenu.Title})
	}

	for _, action := range submenu.actions {
		elements = append(elements, element{key: action.Key, title: action.Title})
	}

	for _, element := range elements {
		text := fmt.Sprintf(" %s- %s ", string(element.key), element.title)
		if len(content[row])+len(text) >= width {
			content = append(content, []byte(text))
			row++
		} else {
			content[row] = append(content[row], []byte(text)...)
		}
	}

	return content
}
