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

