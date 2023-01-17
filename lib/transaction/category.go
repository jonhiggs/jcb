package transaction

import (
	"fmt"
	"jcb/config"
	"jcb/lib/validate"
	"strings"
)

type Category struct {
	value string
}

// Get the string of Category
func (c *Category) GetText() string {
	return (*c).value
}

// Get the string of Category
func (c *Category) GetValue() string {
	return (*c).GetText()
}

// Set the text of Category
func (c *Category) SetText(s string) error {
	s = strings.Trim(s, " ")
	err := validate.Category(s)
	if err != nil {
		return fmt.Errorf("setting category from string: %w", err)
	}

	(*c).value = s
	return nil
}

// Set the string of Category
func (c *Category) SetValue(s string) error {
	return (*c).SetText(s)
}

// To support the Stringer interface
func (c *Category) String() string {
	return fmt.Sprintf("%-*s", config.CATEGORY_MAX_LENGTH, c.value)
}
