package common

func SliceInsertIndex(slice []rune, char rune, index int) []rune {
	tmp := append(slice[:index], append([]rune{char}, slice[index:]...)...)
	return tmp
}
