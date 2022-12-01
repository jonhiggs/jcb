package repeater

import (
	"errors"
	"strconv"
	"time"
)

func Expand(startDate time.Time, endDate time.Time, rule string) ([]time.Time, error) {
	var timestamps []time.Time
	currentYear := startDate.Year()

	u, err := unit(rule)
	if err != nil {
		return timestamps, err
	}

	f, err := frequency(rule)
	if err != nil {
		return timestamps, err
	}

	switch u {
	case "d":
		for i := 0; true; i += f {
			var ts time.Time
			ts = startDate.AddDate(0, 0, i)

			if ts.Unix() <= endDate.Unix() {
				timestamps = append(timestamps, ts)
			} else {
				break
			}
		}

	case "w":
		for i := 0; true; i += f * 7 {
			var ts time.Time
			ts = startDate.AddDate(0, 0, i)

			if ts.Year() == currentYear {
				timestamps = append(timestamps, ts)
			} else {
				break
			}
		}
	case "m":
		for i := 0; true; i += f {
			var ts time.Time
			ts = startDate.AddDate(0, i, 0)

			if ts.Year() == currentYear {
				timestamps = append(timestamps, ts)
			} else {
				break
			}
		}
	}

	return timestamps, nil
}

func frequency(rule string) (int, error) {
	s := rule[0 : len(rule)-1]
	return strconv.Atoi(s)
}

func unit(rule string) (string, error) {
	u := rule[len(rule)-1:]
	if u != "d" && u != "w" && u != "m" {
		return "x", errors.New("Invalid unit of frequency. Expects 'd', 'w' or 'm'.")
	}
	return u, nil
}
