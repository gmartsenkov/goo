package windows

func (window *Window) backspace() {
	cursor := window.Cursor
	slice := window.Content

	x := cursor.X
	y := cursor.Y

	if x == 0 {
		window.SetCursor((len(slice[y-1])), y-1)
		slice[y-1] = append(slice[y-1], slice[y]...)
		slice[y] = []byte{}
		return
	}

	slice[y] = append(slice[y][:x-1], slice[y][x:]...)

	window.MoveCursorLeft()
}

func (window *Window) enter() {
	cursor := window.Cursor
	y := cursor.Y
	x := cursor.X

	slice := window.Content
	tempSlice := [][]byte{}
	tempSlice = append(tempSlice, slice[:y]...)
	tempSlice = append(tempSlice, append([]byte{}, slice[y][:x]...))
	tempSlice = append(tempSlice, append([]byte{}, slice[y][x:]...))
	tempSlice = append(tempSlice, slice[y+1:]...)

	window.Content = tempSlice

	window.SetCursor(0, y+1)
}
