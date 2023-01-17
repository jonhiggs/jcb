// Valiate that data is correct for inserting into the database.
package validate

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Description(s string) error {
	return nil
}

func Date(s string) error {
	splitDate := strings.Split(strings.Trim(s, " "), "-")
	year, _ := strconv.Atoi(splitDate[0])
	month, _ := strconv.Atoi(splitDate[1])
	day, _ := strconv.Atoi(splitDate[2])

	s = fmt.Sprintf("%04d-%02d-%02d", year, month, day)

	_, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("validating date: %w", err)
	}

	return nil
}
