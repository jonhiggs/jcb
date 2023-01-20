package transaction

import (
	"testing"

	"github.com/mcuadros/go-defaults"
)

func TestNew(t *testing.T) {
	d := new(Transaction)
	defaults.SetDefaults(d)

	if d.Id != -1 {
		t.Errorf("got %d, expected -1", d.Id)
	}
}
