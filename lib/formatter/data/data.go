// convert strings to data

package dataFormatter

import (
	"fmt"
	"jcb/domain"
	"jcb/lib/validator"
	"log"
	"strconv"
	"strings"
	"time"
)

func Cents(s string) int64 {
	if validator.Cents(s) != nil {
		log.Fatal("cannot convert invalid input to data")
	}

	s = strings.Trim(s, " ")

	if len(strings.Split(s, ".")) == 1 {
		s = fmt.Sprintf("%s.00", s)
	}

	i, _ := strconv.ParseInt(strings.Replace(s, ".", "", 1), 10, 64)
	return i
}

func Date(s string) time.Time {
	if validator.Date(s) != nil {
		log.Fatal("cannot convert invalid input to data")
	}

	r, _ := time.Parse("2006-01-02", strings.Trim(s, " "))
	return r
}

func Description(s string) string {
	if validator.Description(s) != nil {
		log.Fatal("cannot convert invalid input to data")
	}

	return strings.Trim(s, " ")
}

func Id(d string) int64 {
	if validator.Id(d) != nil {
		log.Fatal("cannot convert invalid input to data")
	}

	id, _ := strconv.ParseInt(d, 10, 64)
	return id
}

func RepeatRule(rule string) string {
	if validator.RepeatRule(rule) != nil {
		log.Fatal("cannot convert invalid input to data")
	}

	return strings.Trim(rule, " ")
}

func RepeatRuleUnit(rule string) string {
	if validator.RepeatRule(rule) != nil {
		log.Fatal("cannot convert invalid input to data")
	}

	return rule[len(rule)-1:]
}

func RepeatRuleFrequency(rule string) int {
	if validator.RepeatRule(rule) != nil {
		log.Fatal("cannot convert invalid input to data")
	}

	i, _ := strconv.Atoi(rule[0 : len(rule)-1])
	return i
}

func Transaction(d domain.StringTransaction) domain.Transaction {
	return domain.Transaction{
		Id(d.Id),
		Date(d.Date),
		Description(d.Description),
		Cents(d.Cents),
	}
}
