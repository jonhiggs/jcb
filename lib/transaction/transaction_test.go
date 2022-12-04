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

func TestBalanceSetD(t *testing.T) {
	set := make([]domain.Transaction, 6)
	set[0] = domain.Transaction{1, time.Now(), "one", 20}
	set[1] = domain.Transaction{1, time.Now(), "two", 20}
	set[2] = domain.Transaction{1, time.Now(), "three", 20}
	set[3] = domain.Transaction{1, time.Now(), "one", 10}
	set[4] = domain.Transaction{1, time.Now(), "two", 10}
	set[5] = domain.Transaction{1, time.Now(), "three", 10}

	got := balanceSet(set, 100)

	if got[0] != 30 {
		t.Error(fmt.Sprintf("expected 30 but got %d", got[0]))
	}

	if got[1] != 50 {
		t.Error(fmt.Sprintf("expected 50 but got %d", got[1]))
	}

	if got[2] != 70 {
		t.Error(fmt.Sprintf("expected 70 but got %d", got[2]))
	}

	if got[3] != 80 {
		t.Error(fmt.Sprintf("expected 80 but got %d", got[3]))
	}

	if got[4] != 90 {
		t.Error(fmt.Sprintf("expected 90 but got %d", got[4]))
	}

	if got[5] != 100 {
		t.Error(fmt.Sprintf("expected 100 but got %d", got[5]))
	}
}
