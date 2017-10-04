package predefined_windows

import "goo/windows"

func Menu() windows.Window {
	window := windows.Window{}
	window.EnableBoldContent = true
	window.EnableBorder = true
	window.EnableSolidForeground = true

	return window
}
