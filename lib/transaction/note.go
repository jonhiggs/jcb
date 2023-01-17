package transaction

import (
	"fmt"
	"jcb/config"
	"jcb/lib/validate"
	"strings"
)

type Note struct {
	value string `default:`
}

// Get the string of Note
func (n *Note) GetText() string {
	return (*n).value
}

// Get the string of Note
func (n *Note) GetValue() string {
	return (*n).GetText()
}

// Set the text of Note
func (n *Note) SetText(s string) error {
	s = strings.Trim(s, " ")
	err := validate.Note(s)
	if err != nil {
		return fmt.Errorf("setting note from string: %w", err)
	}

	(*n).value = s
	return nil
}

// Set the string of Note
func (n *Note) SetValue(s string) error {
	return (*n).SetText(s)
}

// To support the Stringer interface
func (n *Note) String() string {
	return fmt.Sprintf("%-*s", config.NOTE_MAX_LENGTH, n.value)
}

// Returns true if a note exists
func (n *Note) Exists() bool {
	return len((*n).value) != 0
}
