package promptBindings

import (
	"strings"

	"code.rocketnine.space/tslocum/cview"
)

// XXX: not a ring, but one day it might be
var killRing string

// returns the deleted character
func DeleteChar(field *cview.InputField) string {
	pos := field.GetCursorPosition()
	text := field.GetText()
	var deletedChar string

	textSlice := strings.Split(text, "")

	var newSlice []string
	for i, l := range textSlice {
		if i == pos {
			deletedChar = l
			continue
		}

		newSlice = append(newSlice, l)
	}

	field.SetText(strings.Join(newSlice, ""))
	field.SetCursorPosition(pos)
	return deletedChar
}

func BackwardChar(field *cview.InputField) {
	pos := field.GetCursorPosition()
	if pos > 0 {
		field.SetCursorPosition(pos - 1)
	}
}

func ForwardChar(field *cview.InputField) {
	pos := field.GetCursorPosition()
	if pos < len(field.GetText()) {
		field.SetCursorPosition(pos + 1)
	}
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

	var yankString string

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
		yankString += DeleteChar(field)
		pos -= 1
	}
	killRingInsert(reverse(yankString))
}

func Yank(field *cview.InputField) {
	pos := field.GetCursorPosition()
	text := field.GetText()

	var newText string

	for i, l := range strings.Split(text, "") {
		if i == pos {
			newText += killRing
		}

		newText += l
	}

	if pos == len(text) {
		newText += killRing
	}

	field.SetText(newText)
	field.SetCursorPosition(pos + len(killRing))
}

func killRingInsert(s string) {
	killRing = s
}

func reverse(s string) string {
	var newString string
	for i := len(s) - 1; i >= 0; i-- {
		newString = newString + string(s[i])
	}
	return newString
}
