// read fields and validate their content.

package fieldReader

import (
	"errors"
	"fmt"
	"strings"

	gc "github.com/rthornton128/goncurses"
)

func AsAmount(field *gc.Field) (string, error) {
	str := strings.Trim(field.Buffer(), " ")
	amountSplit := strings.Split(str, ".")
	if len(amountSplit) > 2 {
		return "", errors.New("Amount has too many dots")
	}
	if len(amountSplit) == 2 && len(amountSplit[1]) > 2 {
		return "", errors.New(fmt.Sprintf("Amount has to many decimal places [%d]", len(amountSplit[1])))
	}

	return str, nil
}
