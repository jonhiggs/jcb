// read fields and validate their content.

package uiFieldReader

import (
	"errors"
	"fmt"
	"strings"

	gc "github.com/rthornton128/goncurses"
)

func Amount(field *gc.Field) (string, error) {
	str := strings.Trim(field.Buffer(), " ")
	amountSplit := strings.Split(amountStr, ".")
	if len(amountSplit) > 2 {
		return -1, errors.New("Amount has too many dots")
	}
	if len(amountSplit) == 2 && len(amountSplit[1]) > 2 {
		return -1, errors.New(fmt.Sprintf("Amount has to many decimal places [%d]", len(amountSplit[1])))
	}

	//t := unformatTransaction(FormattedTransaction{idStr, dateStr, descriptionStr, amountStr})
}
