// convert a string from an input box and format into data
package format

import (
	"fmt"
	"jcb/lib/validate"
	"strconv"
	"strings"
	"time"
)

// convert a description string into data
func Description(s string) (string, error) {
	s = strings.Trim(s, " ")
	err := validate.Description(s)
	if err != nil {
		return s, fmt.Errorf("formatting description: %w", err)
	}

	return s, nil
}

// convert a description into a string
func DescriptionString(d string) string {
	return strings.Trim(d, " ")
}

// convert a date string into data
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

// convert a date into a string
func DateString(d time.Time) string {
	return d.Format("2006-01-02")
}

// return cents as a string
func CentsString(i int) string {
	var d string
	var c string

	negative := ""
	if i < 0 {
		negative = "-"
		i = i * -1
	}

	s := fmt.Sprintf("%d", i)

	if len(s) == 2 {
		d = "0"
		c = s
	} else if len(s) == 1 {
		d = "0"
		c = fmt.Sprintf("0%s", s)
	} else {
		d = s[0 : len(s)-2]
		c = s[len(s)-2 : len(s)]
	}

	return fmt.Sprintf("%s%s.%s", negative, d, c)
}
