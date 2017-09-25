package editors

const (
	StateNormal = 0
	StateInsert = 1
	StateVisual = 2
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
