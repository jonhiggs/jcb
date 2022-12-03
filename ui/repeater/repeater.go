package repeater

import (
	dataf "jcb/ui/formatter/data"
	"time"
)

func Expand(date time.Time, startDate time.Time, endDate time.Time, rule string) ([]time.Time, error) {
	var timestamps []time.Time

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
			ts = date.AddDate(0, 0, i)

			if ts.Unix() < startDate.Unix() {
				continue
			} else if ts.Unix() <= endDate.Unix() {
				timestamps = append(timestamps, ts)
			} else {
				break
			}

			if f == 0 {
				break
			}
		}

	case "w":
		for i := 0; true; i += f * 7 {
			var ts time.Time
			ts = date.AddDate(0, 0, i)

			if ts.Unix() < startDate.Unix() {
				continue
			} else if ts.Unix() <= endDate.Unix() {
				timestamps = append(timestamps, ts)
			} else {
				break
			}
		}
	case "m":
		for i := 0; true; i += f {
			var ts time.Time
			ts = date.AddDate(0, i, 0)

			if ts.Unix() < startDate.Unix() {
				continue
			} else if ts.Unix() <= endDate.Unix() {
				timestamps = append(timestamps, ts)
			} else {
				break
			}
		}
	}

	return timestamps, nil
}
