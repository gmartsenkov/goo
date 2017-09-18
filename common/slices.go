package common

func SliceInsertIndex(slice []byte, char byte, index int) []byte {
	tmp := append(slice[:index], append([]byte{char}, slice[index:]...)...)
	return tmp
}
