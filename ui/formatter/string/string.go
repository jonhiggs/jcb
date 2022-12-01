// convert data to strings

package stringFormatter

import (
	"errors"
	"fmt"
	"jcb/domain"
	"strconv"
	"strings"
	"time"
)

func Cents(i int64) (string, error) {
	var d string
	var c string

	negative := ""
	if i < 0 {
		negative = "-"
		i = i * -1
	}

	s := fmt.Sprintf("%d", i)

	if len(s) == 2 {
		d = "0"
		c = s
	} else if len(s) == 1 {
		d = "0"
		c = fmt.Sprintf("0%s", s)
	} else {
		d = s[0 : len(s)-2]
		c = s[len(s)-2 : len(s)]
	}
	return fmt.Sprintf("%s%s.%s", negative, d, c), nil
}

func Date(d time.Time) (string, error) {
	year, err := strconv.ParseInt(d.Format("2006"), 10, 64)
	if err != nil || year < 2000 {
		return "", errors.New("Invalid date")
	}
	return d.Format("2006-01-02"), nil
}

func Description(d string) (string, error) {
	return strings.Trim(d, " "), nil
}

func Id(d int64) (string, error) {
	var err error
	err = nil

	s := fmt.Sprintf("%d", d)

	if d < 0 {
		s = "0"
		err = errors.New("Id cannot be less than 0")
	}
	return s, err
}

func Transaction(d domain.Transaction) (domain.StringTransaction, error) {
	r := domain.StringTransaction{}
	id, err := Id(d.Id)
	if err != nil {
		return r, err
	}

	date, err := Date(d.Date)
	if err != nil {
		return r, err
	}

	description, err := Description(d.Description)
	if err != nil {
		return r, err
	}

	cents, err := Cents(d.Cents)
	if err != nil {
		return r, err
	}

	return domain.StringTransaction{id, date, description, cents}, nil
}
