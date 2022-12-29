package acceptanceFunction

import (
	"jcb/config"
	"log"
	"regexp"

	"code.rocketnine.space/tslocum/cview"
)

func Description(textToCheck string, lastChar rune) bool {
	return cview.InputFieldMaxLength(config.DESCRIPTION_MAX_LENGTH)(textToCheck, lastChar)
}

func Notes(textToCheck string, lastChar rune) bool {
	return cview.InputFieldMaxLength(config.NOTES_MAX_LENGTH)(textToCheck, lastChar)
}

func Category(textToCheck string, lastChar rune) bool {
	return cview.InputFieldMaxLength(config.CATEGORY_MAX_LENGTH)(textToCheck, lastChar)
}

func Date(textToCheck string, lastChar rune) bool {
	if len(textToCheck) > 10 {
		return false
	}

	valid_char, err := regexp.MatchString(`^[0-9\-]*$`, textToCheck)
	if err != nil {
		log.Fatal(err)
	}
	return valid_char
}

func Cents(textToCheck string, lastChar rune) bool {
	if len(textToCheck) > 10 {
		return false
	}

	valid_char, err := regexp.MatchString(`^\-?[0-9]*(\.[0-9]{0,2})?$`, textToCheck)
	if err != nil {
		log.Fatal(err)
	}
	return valid_char
}

func Any(textToCheck string, lastChar rune) bool {
	return true
}
