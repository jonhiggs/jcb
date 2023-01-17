package validator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Id(d string) error {
	id, _ := strconv.ParseInt(d, 10, 64)
	if id < 0 {
		return errors.New("Id cannot be less than 0")
	}
	return nil
}

func RepeatRule(rule string) error {
	rule = strings.Trim(rule, " ")
	re := regexp.MustCompile(`^[0-9]+[dwmy]$`)

	if !re.MatchString(rule) {
		return errors.New("Invalid repeat rule")
	}

	return nil
}
