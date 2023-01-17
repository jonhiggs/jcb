// convert strings to data

package dataFormatter

import (
	"fmt"
	"jcb/lib/validator"
	"log"
	"strconv"
	"strings"
)

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
