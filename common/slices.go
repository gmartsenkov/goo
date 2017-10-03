package common

func SliceInsertIndex(slice []Cell, char rune, index int) []Cell {
	tmp := append(slice[:index], append([]Cell{Cell{Ch: char}}, slice[index:]...)...)
	return tmp
}
