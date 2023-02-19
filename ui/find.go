package ui

import "regexp"

var query *regexp.Regexp

func setQuery(q string) error {
	re, err := regexp.Compile(q)
	if err != nil {
		return err
	}

	query = re
	return nil
}
