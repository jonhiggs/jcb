// convert data from string to data

package dataFormatter

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Cents(s string) int64 {
	cents, _ := strconv.ParseInt(strings.Replace(strings.Trim(s, " "), ".", "", 1), 10, 64)
	return cents
}

func Date(s string) time.Time {
	time, _ := time.Parse("2006-01-02", strings.Trim(s, " "))
	return time
}

func Description(s string) string {
	return strings.Trim(s, " ")
}

func Id(d int64) string {
	return fmt.Sprintf("%d", d)
}
