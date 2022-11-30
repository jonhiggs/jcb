// convert data to strings

package dataFormatter

import (
	"jcb/domain"
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

func Id(d string) int64 {
	r, _ := strconv.ParseInt(d, 10, 64)
	return r
}

func Transaction(d domain.Transaction) domain.StringTransaction {
	return domain.Transaction{
		Id(d.Id),
		Date(d.Date),
		Description(d.Description),
		Cents(d.Cents),
	}
}
