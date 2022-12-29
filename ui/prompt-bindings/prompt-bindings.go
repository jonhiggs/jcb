package promptBindings

import (
	"strings"

	"code.rocketnine.space/tslocum/cview"
)

func DeleteChar(field *cview.InputField) {
	pos := field.GetCursorPosition()
	text := field.GetText()

	textSlice := strings.Split(text, "")

	var newSlice []string
	for i, l := range textSlice {
		if i == pos {
			continue
		}

		newSlice = append(newSlice, l)
	}

	field.SetText(strings.Join(newSlice, ""))
	field.SetCursorPosition(pos)
}

func BackwardChar(field *cview.InputField) {
	pos := field.GetCursorPosition()
	field.SetCursorPosition(pos - 1)
}

func ForwardChar(field *cview.InputField) {
	pos := field.GetCursorPosition()
	field.SetCursorPosition(pos + 1)
}

func UnixWordRubout(field *cview.InputField) {
	pos := field.GetCursorPosition()
	i := 0

	for pos > 0 {
		if field.GetText()[pos-1] == ' ' && i > 0 {
			break
		}

		// delete all the spaces before considering anything deleted
		if field.GetText()[pos-1] != ' ' {
			i += 1
		}

		BackwardChar(field)
		DeleteChar(field)
		pos -= 1
	}
}

func OtherUnixWordRubout(field *cview.InputField) {
	pos := field.GetCursorPosition()
	i := 0

	for pos > 0 {
		if (field.GetText()[pos-1] == ' ' || field.GetText()[pos-1] == '-') && i > 0 {
			break
		}

		// delete all the spaces before considering anything deleted
		if field.GetText()[pos-1] != ' ' {
			i += 1
		}

		BackwardChar(field)
		DeleteChar(field)
		pos -= 1
	}
}
