package transaction

import (
	"testing"
)

func TestNew(t *testing.T) {
	d := NewTransaction()

	if d.Id != -1 {
		t.Errorf("got %d, expected -1", d.Id)
	}

	if d.Committed {
		t.Errorf("got %v, expected %v", d.Committed, false)
	}
}
