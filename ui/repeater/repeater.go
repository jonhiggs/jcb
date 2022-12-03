package repeater

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func Expand(date time.Time, startDate time.Time, endDate time.Time, rule string) ([]time.Time, error) {
	if strings.HasPrefix(rule, "0") {
		if date.Unix() >= startDate.Unix() && date.Unix() < endDate.Unix() {
			return []time.Time{date}, nil
		} else {
			return []time.Time{}, nil
		}
	} else {
		timestamps := []time.Time{date}
		for date.Unix() >= startDate.Unix() && date.Unix() < endDate.Unix() {
			date, err := relativeDate(date, rule)
			if err != nil {
				return timestamps, err
			}
			timestamps = append(timestamps, date)
		}
		return timestamps, nil
	}
}

func relativeDate(date time.Time, rule string) (time.Time, error) {
	f, err := strconv.Atoi(rule[0 : len(rule)-1])
	if err != nil {
		return date, err
	}
	u := rule[len(rule)-1 : len(rule)]

	if f == 0 {
		return date, errors.New("Pattern doesn't repeat")
	}

	switch u {
	case "d":
		return date.AddDate(0, 0, 1*f), nil
	case "w":
		return date.AddDate(0, 0, 7*f), nil
	case "m":
		return date.AddDate(0, 1, 0*f), nil
	}

	return date, errors.New("Shouldn't have got here.")
}
