package repeater

import (
	"jcb/lib/dates"
	"strconv"
	"strings"
	"time"
)

func Expand(date time.Time, endDate time.Time, rule string) ([]time.Time, error) {
	if strings.HasPrefix(rule, "0") {
		if date.Unix() < dates.LastCommitted().Unix() {
			return []time.Time{}, nil
		}

		t, err := relativeDate(date, rule, 1)
		return []time.Time{t}, err
	}

	var timestamps []time.Time
	for i := 0; i < 380; i++ {
		nextDate, _ := relativeDate(date, rule, i)

		if nextDate.Unix() == date.Unix() {
			continue
		}

		if nextDate.Unix() < dates.LastCommitted().Unix() {
			continue
		}

		if nextDate.Unix() >= endDate.Unix() {
			break
		}

		timestamps = append(timestamps, nextDate)
	}

	return timestamps, nil
}

func relativeDate(date time.Time, rule string, multiplier int) (time.Time, error) {
	f, err := strconv.Atoi(rule[0 : len(rule)-1])
	if err != nil {
		return date, err
	}
	u := rule[len(rule)-1 : len(rule)]

	switch u {
	case "d":
		date = date.AddDate(0, 0, (multiplier * 1 * f))
	case "w":
		date = date.AddDate(0, 0, (multiplier * 1 * (f * 7)))
	case "m":
		date = date.AddDate(0, (multiplier * 1 * f), 0)
	case "y":
		date = date.AddDate((multiplier * 1 * f), 0, 0)
	}

	return date, nil
}
