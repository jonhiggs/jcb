package repeater

import (
	"errors"
	dataf "jcb/ui/formatter/data"
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
			date, err := nextDate(date, rule)
			if err != nil {
				return timestamps, err
			}
			timestamps = append(timestamps, date)
		}
		return timestamps, nil
	}
}

func nextDate(date time.Time, rule string) (time.Time, error) {
	u, err := dataf.RepeatRuleUnit(rule)
	if err != nil {
		return date, err
	}

	f, err := dataf.RepeatRuleFrequency(rule)
	if err != nil {
		return date, err
	}

	if f == 0 {
		return date, errors.New("Pattern doesn't repeat")
	}

	switch u {
	case "d":
		return date.AddDate(0, 0, 1), nil
	case "w":
		return date.AddDate(0, 0, 7), nil
	case "m":
		return date.AddDate(0, 1, 0), nil
	}

	return date, errors.New("Shouldn't have got here.")
}
