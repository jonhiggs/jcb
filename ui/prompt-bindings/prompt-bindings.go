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
	separators := []rune{' '}
	deleteBackwardsWithSeparators(field, separators)
}

func OtherUnixWordRubout(field *cview.InputField) {
	separators := []rune{' ', '-', '.', '/'}
	deleteBackwardsWithSeparators(field, separators)
}

func deleteBackwardsWithSeparators(field *cview.InputField, separators []rune) {
	pos := field.GetCursorPosition()
	i := 0

all:
	for pos > 0 {
		if i > 0 {
			for _, s := range separators {
				if field.GetText()[pos-1] == byte(s) {
					break all
				}
			}
		}

		// delete all the spaces before considering anything deleted
		foundSeparator := false
		for _, s := range separators {
			if field.GetText()[pos-1] == byte(s) {
				foundSeparator = true
			}
		}
		if !foundSeparator {
			i += 1
		}

		BackwardChar(field)
		DeleteChar(field)
		pos -= 1
	}
}
