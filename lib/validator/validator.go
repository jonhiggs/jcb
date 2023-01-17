package validator

import (
	"errors"
	"regexp"
	"strings"
)

func RepeatRule(rule string) error {
	rule = strings.Trim(rule, " ")
	re := regexp.MustCompile(`^[0-9]+[dwmy]$`)

	if !re.MatchString(rule) {
		return errors.New("Invalid repeat rule")
	}

	return nil
}
