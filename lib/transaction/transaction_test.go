package transaction

import (
	"fmt"
	"jcb/domain"
	"testing"
	"time"
)

func TestBalanceSetA(t *testing.T) {
	set := make([]domain.Transaction, 1)
	set[0] = domain.Transaction{1, time.Now(), "one", 100}

	got := balanceSet(set, 100)

	if got[0] != 100 {
		t.Error(fmt.Sprintf("expected 100 but got %d", got[0]))
	}

	if len(got) != 1 {
		t.Error(fmt.Sprintf("expected 1 element but got %d", len(got)))
	}
}

func TestBalanceSetB(t *testing.T) {
	set := make([]domain.Transaction, 2)
	set[0] = domain.Transaction{1, time.Now(), "one", 100}
	set[1] = domain.Transaction{1, time.Now(), "two", 100}

	got := balanceSet(set, 100)

	if got[0] != 0 {
		t.Error(fmt.Sprintf("expected 0 but got %d", got[0]))
	}

	if got[1] != 100 {
		t.Error(fmt.Sprintf("expected 100 but got %d", got[1]))
	}

	if len(got) != 2 {
		t.Error(fmt.Sprintf("expected 1 element but got %d", len(got)))
	}
}

func TestBalanceSetC(t *testing.T) {
	set := make([]domain.Transaction, 3)
	set[0] = domain.Transaction{1, time.Now(), "one", -80}
	set[1] = domain.Transaction{1, time.Now(), "two", -100}
	set[2] = domain.Transaction{1, time.Now(), "three", 1000}

	got := balanceSet(set, 1000)

	if got[0] != 1180 {
		t.Error(fmt.Sprintf("expected 1180 but got %d", got[0]))
	}

	if got[1] != 1100 {
		t.Error(fmt.Sprintf("expected 1100 but got %d", got[1]))
	}

	if got[2] != 1000 {
		t.Error(fmt.Sprintf("expected 1000 but got %d", got[2]))
	}

	if len(got) != 3 {
		t.Error(fmt.Sprintf("expected 1 element but got %d", len(got)))
	}
}
