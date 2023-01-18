package transaction

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Cents struct {
	value int
}

// Get the cents.
func (c *Cents) GetValue() int { return (*c).value }

// Set the cents.
func (c *Cents) SetValue(i int) error {
	(*c).value = i
	return nil
}

// Get cents as a string
func (c *Cents) GetText() string {
	i := (*c).value
	var dollars string
	var cents string

	negative := ""
	if i < 0 {
		negative = "-"
		i = i * -1
	}

	s := fmt.Sprintf("%d", i)

	if len(s) == 2 {
		dollars = "0"
		cents = s
	} else if len(s) == 1 {
		dollars = "0"
		cents = fmt.Sprintf("0%s", s)
	} else {
		dollars = s[0 : len(s)-2]
		cents = s[len(s)-2 : len(s)]
	}

	return fmt.Sprintf("%s%s.%s", negative, dollars, cents)
}

// Set the cents from a string.
func (c *Cents) SetText(s string) error {
	if !ValidCentString(s) {
		return fmt.Errorf("invalid cents")
	}

	s = strings.Trim(s, " ")

	if len(strings.Split(s, ".")) == 1 {
		s = fmt.Sprintf("%s.00", s)
	}

	s = strings.Replace(s, ".", "", 1)
	i, _ := strconv.Atoi(s)

	(*c).value = i

	return nil
}

func (c *Cents) String() string {
	return fmt.Sprintf("%10s", c.GetText())
}

func (c *Cents) IsDebit() bool {
	return (*c).value < 0
}

func (c *Cents) IsCredit() bool {
	return !(*c).IsDebit()
}

func (c *Cents) Add(i int) int {
	(*c).value += i
	return (*c).value
}

func ValidCentString(s string) bool {
	s = strings.Trim(s, " ")

	re := regexp.MustCompile(`^-?[0-9]+(\.[0-9]{1,2})?$`)
	if !re.MatchString(s) {
		return false
	}

	return true
}
