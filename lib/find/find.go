package find

import (
	"jcb/config"
	"regexp"
	"strings"

	"code.rocketnine.space/tslocum/cview"
)

var query *regexp.Regexp

func SetQuery(q string) error {
	re, err := regexp.Compile(q)
	if err != nil {
		return err
	}

	query = re
	return nil
}

func TableRowMatches(table *cview.Table, row int) bool {
	if query == nil {
		return false
	}

	if row == 0 {
		return false
	}

	category := strings.Trim(table.GetCell(row, config.CATEGORY_COLUMN).GetText(), " ")
	description := strings.Trim(table.GetCell(row, config.DESCRIPTION_COLUMN).GetText(), " ")

	if query.MatchString(category) {
		return true
	}

	if query.MatchString(description) {
		return true
	}

	return false
}
