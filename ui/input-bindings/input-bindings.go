package inputBindings

import (
	"strings"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

// XXX: not a ring, but one day it might be
var killRing string

func Configuration(handler func(ev *tcell.EventKey) *tcell.EventKey) *cbind.Configuration {
	c := cbind.NewConfiguration()
	c.SetKey(tcell.ModCtrl, tcell.KeyCtrlD, handler)
	c.SetKey(tcell.ModCtrl, tcell.KeyCtrlF, handler)
	c.SetKey(tcell.ModCtrl, tcell.KeyCtrlB, handler)
	c.SetKey(tcell.ModCtrl, tcell.KeyCtrlK, handler)
	c.SetKey(tcell.ModCtrl, tcell.KeyCtrlW, handler)
	c.SetKey(tcell.ModCtrl, tcell.KeyCtrlY, handler)
	c.SetRune(tcell.ModAlt, 'd', handler)
	c.SetRune(tcell.ModAlt, 'f', handler)
	c.SetRune(tcell.ModAlt, 'b', handler)
	c.SetKey(tcell.ModAlt, tcell.KeyBackspace2, handler)

	for _, k := range []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*(),./<>?;':\"[]{}-+") {
		c.SetRune(0, k, handler)
	}

	return c
}

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

func BackwardWord(field *cview.InputField) {
	separators := []rune{' ', '-', '.', '/'}
	foundChar := false

all:
	for field.GetCursorPosition() > 0 {
		char := field.GetText()[field.GetCursorPosition()-1]
		for _, s := range separators {
			if char == byte(s) {
				if foundChar {
					return
				} else {
					BackwardChar(field)
					continue all
				}
			}
		}

		foundChar = true
		BackwardChar(field)
	}
}

func ForwardChar(field *cview.InputField) {
	pos := field.GetCursorPosition()
	if pos < len(field.GetText()) {
		field.SetCursorPosition(pos + 1)
	}
}

func ForwardWord(field *cview.InputField) {
	separators := []rune{' ', '-', '.', '/'}
	foundChar := false

all:
	for field.GetCursorPosition() < len(field.GetText()) {
		char := field.GetText()[field.GetCursorPosition()]
		for _, s := range separators {
			if char == byte(s) {
				if foundChar {
					return
				} else {
					ForwardChar(field)
					continue all
				}
			}
		}

		foundChar = true
		ForwardChar(field)
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

func DeleteWord(field *cview.InputField) {
	separators := []rune{' ', '-', '.', '/'}
	deleteForwardsWithSeparators(field, separators)
}

func KillLine(field *cview.InputField) {
	pos := field.GetCursorPosition()
	text := field.GetText()
	killRing = text[pos:len(text)]
	field.SetText(text[0:pos])
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

func deleteForwardsWithSeparators(field *cview.InputField, separators []rune) {
	pos := field.GetCursorPosition()
	i := 0

	var yankString string

all:
	for pos < len(field.GetText()) {
		if i > 0 {
			for _, s := range separators {
				if field.GetText()[pos] == byte(s) {
					break all
				}
			}
		}

		// delete all the spaces before considering anything deleted
		foundSeparator := false
		for _, s := range separators {
			if field.GetText()[pos] == byte(s) {
				foundSeparator = true
			}
		}
		if !foundSeparator {
			i += 1
		}

		yankString += DeleteChar(field)
	}
	killRingInsert(yankString)
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
