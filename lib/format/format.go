// convert a string from an input box and format into data
package format

import (
	"fmt"
	"jcb/lib/validate"
	"strconv"
	"strings"
	"time"
)

func Description(s string) (string, error) {
	s = strings.Trim(s, " ")
	err := validate.Description(s)
	if err != nil {
		return s, fmt.Errorf("formatting description: %w", err)
	}

	return s, nil
}

func Date(s string) (time.Time, error) {
	splitDate := strings.Split(strings.Trim(s, " "), "-")
	year, _ := strconv.Atoi(splitDate[0])
	month, _ := strconv.Atoi(splitDate[1])
	day, _ := strconv.Atoi(splitDate[2])

	s = fmt.Sprintf("%04d-%02d-%02d", year, month, day)

	err := validate.Date(s)
	if err != nil {
		return time.Unix(0, 0), fmt.Errorf("formatting date: %w", err)
	}

	r, _ := time.Parse("2006-01-02", s)
	return r, nil
}
