package repeater

import (
	dataf "jcb/ui/formatter/data"
	"time"
)

func Expand(startDate time.Time, endDate time.Time, rule string) ([]time.Time, error) {
	var timestamps []time.Time
	currentYear := startDate.Year()

	u, err := dataf.RepeatRuleUnit(rule)
	if err != nil {
		return timestamps, err
	}

	f, err := dataf.RepeatRuleFrequency(rule)
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
