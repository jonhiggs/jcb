package acceptanceFunction

import (
	"jcb/config"
	"log"
	"regexp"
	"strings"

	"code.rocketnine.space/tslocum/cview"
)

func FieldFunc(field *cview.InputField) func(field *cview.InputField) bool {
	switch field.GetLabel() {
	case "Date:":
		return Date
	case "Description:":
		return Description
	case "Category:":
		return Category
	case "Amount:":
		return Cents
	case "Notes:":
		return Notes
	default:
		return Any
	}
}

func Description(field *cview.InputField) bool {
	if len(field.GetText()) > config.DESCRIPTION_MAX_LENGTH {
		return false
	}
	return true
}

func Notes(field *cview.InputField) bool {
	if len(field.GetText()) > config.NOTE_MAX_LENGTH {
		return false
	}
	return true
}

func Category(field *cview.InputField) bool {
	if len(field.GetText()) > config.CATEGORY_MAX_LENGTH {
		return false
	}

	if len(strings.Fields(field.GetText())) > 1 {
		return false
	}

	return true
}

func Date(field *cview.InputField) bool {
	if len(field.GetText()) > 10 {
		return false
	}

	valid, err := regexp.MatchString(`^[0-9\-]*$`, field.GetText())
	if err != nil {
		log.Fatal(err)
	}
	return valid
}

func Cents(field *cview.InputField) bool {
	if len(field.GetText()) > 10 {
		return false
	}

	valid, err := regexp.MatchString(`^\-?[0-9]*(\.[0-9]{0,2})?$`, field.GetText())
	if err != nil {
		log.Fatal(err)
	}
	return valid
}

func Any(field *cview.InputField) bool {
	return true
}
