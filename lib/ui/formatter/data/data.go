// convert strings to data

package dataFormatter

import (
	"errors"
	"jcb/domain"
	"strconv"
	"strings"
	"time"
)

func Cents(s string) (int64, error) {
	return strconv.ParseInt(strings.Replace(strings.Trim(s, " "), ".", "", 1), 10, 64)
}

func Date(s string) (time.Time, error) {
	return time.Parse("2006-01-02", strings.Trim(s, " "))
}

func Description(s string) (string, error) {
	return strings.Trim(s, " "), nil
}

func Id(d string) (int64, error) {
	id, err := strconv.ParseInt(d, 10, 64)
	if id < 0 {
		id = 0
		err = errors.New("Id cannot be less than 0")
	}
	return id, err
}

func Transaction(d domain.StringTransaction) (domain.Transaction, error) {
	r := domain.Transaction{}
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

	return domain.Transaction{id, date, description, cents}, nil
}
