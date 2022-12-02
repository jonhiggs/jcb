// convert strings to data

package dataFormatter

import (
	"errors"
	"fmt"
	"jcb/domain"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Cents(s string) (int64, error) {
	s = strings.Trim(s, " ")
	amountSplit := strings.Split(s, ".")
	if len(amountSplit) > 2 {
		return 0, errors.New("Amount has too many dots")
	}
	if len(amountSplit) == 2 && len(amountSplit[1]) > 2 {
		return 0, errors.New(fmt.Sprintf("Amount has to many decimal places [%d]", len(amountSplit[1])))
	}
	if len(amountSplit) == 1 {
		s = fmt.Sprintf("%s.00", s)
	}

	return strconv.ParseInt(strings.Replace(s, ".", "", 1), 10, 64)
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

func RepeatRule(rule string) (string, error) {
	rule = strings.Trim(rule, " ")
	re := regexp.MustCompile(`^[0-9]+[dwm]$`)

	if re.MatchString(rule) {
		return rule, nil
	} else {
		return rule, errors.New("Invalid rule")
	}

}

func RepeatRuleUnit(rule string) (string, error) {
	_, err := RepeatRule(rule)
	if err != nil {
		return rule, err
	}

	u := rule[len(rule)-1:]
	if u != "d" && u != "w" && u != "m" {
		return "x", errors.New(fmt.Sprintf("Invalid unit of frequency [%s]. Expects 'd', 'w' or 'm'.", rule))
	}
	return u, nil
}

func RepeatRuleFrequency(rule string) (int, error) {
	_, err := RepeatRule(rule)
	if err != nil {
		return -1, err
	}

	s := rule[0 : len(rule)-1]
	return strconv.Atoi(s)
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