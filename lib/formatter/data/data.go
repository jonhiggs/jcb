// convert strings to data

package dataFormatter

import (
	"fmt"
	"jcb/lib/validator"
	"log"
	"strconv"
	"strings"
)

func Cents(s string) int64 {
	if validator.Cents(s) != nil {
		log.Fatal(fmt.Sprintf("cannot convert invalid cents '%s' to data", s))
	}

	s = strings.Trim(s, " ")

	if len(strings.Split(s, ".")) == 1 {
		s = fmt.Sprintf("%s.00", s)
	}

	i, _ := strconv.ParseInt(strings.Replace(s, ".", "", 1), 10, 64)
	return i
}

func Category(s string) string {
	if validator.Category(s) != nil {
		log.Fatal("cannot convert invalid cateogry to data")
	}

	return strings.Trim(s, " ")
}

func Notes(s string) string {
	if validator.Notes(s) != nil {
		log.Fatal("cannot convert invalid notes to data")
	}

	return strings.Trim(s, " ")
}

func Id(d string) int64 {
	if validator.Id(d) != nil {
		log.Fatal(fmt.Sprintf("cannot convert invalid id '%s' to data", d))
	}

	id, _ := strconv.ParseInt(d, 10, 64)
	return id
}

func RepeatRule(rule string) string {
	if validator.RepeatRule(rule) != nil {
		log.Fatal(fmt.Sprintf("cannot convert invalid repeat rule '%s' to data", rule))
	}

	return strings.Trim(rule, " ")
}

func RepeatRuleUnit(rule string) string {
	if validator.RepeatRule(rule) != nil {
		log.Fatal("cannot convert invalid repeat rule unit to data")
	}

	return rule[len(rule)-1:]
}

func RepeatRuleFrequency(rule string) int {
	if validator.RepeatRule(rule) != nil {
		log.Fatal("cannot convert invalid frequency to data")
	}

	i, _ := strconv.Atoi(rule[0 : len(rule)-1])
	return i
}
