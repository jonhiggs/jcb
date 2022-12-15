package validator

import (
	"errors"
	"jcb/domain"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Cents(s string) error {
	s = strings.Trim(s, " ")

	re := regexp.MustCompile(`^-?[0-9]+(\.[0-9]{1,2})?$`)
	if !re.MatchString(s) {
		return errors.New("Invalid amount")
	}

	return nil
}

func Date(s string) error {
	_, e := time.Parse("2006-01-02", strings.Trim(s, " "))
	if e != nil {
		return errors.New("Invalid date")
	}
	return nil
}

func Description(s string) error {
	return nil
}

func Id(d string) error {
	id, _ := strconv.ParseInt(d, 10, 64)
	if id < 0 {
		return errors.New("Id cannot be less than 0")
	}
	return nil
}

func RepeatRule(rule string) error {
	rule = strings.Trim(rule, " ")
	re := regexp.MustCompile(`^[0-9]+[dwm]$`)

	if !re.MatchString(rule) {
		return errors.New("Invalid repeat rule")
	}

	return nil
}

func Transaction(d domain.StringTransaction) error {
	e := Id(d.Id)
	if e != nil {
		return e
	}

	e = Date(d.Date)
	if e != nil {
		return e
	}

	e = Description(d.Description)
	if e != nil {
		return e
	}

	e = Cents(d.Cents)
	if e != nil {
		return e
	}

	return nil
}
