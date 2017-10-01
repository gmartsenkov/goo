package editors

const (
	StateNormal uint8 = 0
	StateInsert uint8 = 1
	StateVisual uint8 = 2
)

func (editor *Editor) InsertState() {
	editor.State = StateInsert
}

func (editor *Editor) NormalState() {
	editor.State = StateNormal
}

func (editor *Editor) StateInWords() string {
	switch editor.State {
	case StateNormal:
		return "--NORMAL--"
	case StateInsert:
		return "--INSERT--"
	case StateVisual:
		return "--VISUAL--"
	}

	return ""
}
