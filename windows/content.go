package windows

import (
	"bytes"
	"goo/common"
)

func (window *Window) SplitAndSetContent(content []byte) {
	tmp := bytes.Split(content, []byte("\n"))

	for _, line := range tmp {
		window.Content = append(window.Content, common.BytesToCells(line))
	}
}

func (window *Window) ContentAsRuneArray() [][]rune {
	tmp := [][]rune{}
	for _, line := range window.Content {
		tmp = append(tmp, common.CellsAsRuneArray(line))
	}

	return tmp
}

func (window *Window) AppendLineRuneArray(line []rune) {
	window.Content = append(window.Content, common.RunesToCells(line))
}
