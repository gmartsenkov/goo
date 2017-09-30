package menu

import (
	"errors"
	"goo/editors"
)

type Action struct {
	Title string
	Key   rune
	Fn    func(*editors.Editor)
}

type Actions []Action

func (actions Actions) FindByKey(key rune) (Action, error) {
	for _, action := range actions {
		if action.Key == key {
			return action, nil
		}
	}
	return Action{}, errors.New("ActionNotFound")
}
