// Valiate that data is correct for inserting into the database.
package validate

import (
	"errors"
	"fmt"
	"jcb/config"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Description(s string) error {
	return nil
}

func Date(s string) error {
	splitDate := strings.Split(strings.Trim(s, " "), "-")
	year, _ := strconv.Atoi(splitDate[0])
	month, _ := strconv.Atoi(splitDate[1])
	day, _ := strconv.Atoi(splitDate[2])

	s = fmt.Sprintf("%04d-%02d-%02d", year, month, day)

	_, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("validating date: %w", err)
	}

	return nil
}

func Cents(s string) error {
	s = strings.Trim(s, " ")

	re := regexp.MustCompile(`^-?[0-9]+(\.[0-9]{1,2})?$`)
	if !re.MatchString(s) {
		return errors.New("invalid amount")
	}

	return nil
}

func Note(s string) error {
	return nil
}

func Category(s string) error {
	if len(s) > config.CATEGORY_MAX_LENGTH {
		return errors.New("category too long")
	}

	return nil
}

func RepeatRule(rule string) error {
	rule = strings.Trim(rule, " ")
	re := regexp.MustCompile(`^[0-9]+[dwmy]$`)

	if !re.MatchString(rule) {
		return errors.New("invalid repeat rule")
	}

	return nil
}
