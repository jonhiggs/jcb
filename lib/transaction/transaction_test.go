package transaction

import (
	"fmt"
	"jcb/domain"
	"testing"
	"time"
)

func TestBalanceSetA(t *testing.T) {
	set := make([]domain.Transaction, 1)
	set[0] = domain.Transaction{1, time.Now(), "one", 123, ""}

	got := balanceSet(set, 100)

	if got[0].Id != 1 {
		t.Error(fmt.Sprintf("expected id 1 but got %d", got[0].Id))
	}

	if got[0].Cents != 123 {
		t.Error(fmt.Sprintf("expected cents of 123, got %d", got[0].Cents))
	}

	if got[0].Balance != 100 {
		t.Error(fmt.Sprintf("expected balance of 100, got %d", got[0].Balance))
	}

	if len(got) != 1 {
		t.Error(fmt.Sprintf("expected 1 element, got %d", len(got)))
	}
}

func TestBalanceSetB(t *testing.T) {
	set := make([]domain.Transaction, 2)
	set[0] = domain.Transaction{0, time.Now(), "one", -100, ""}
	set[1] = domain.Transaction{1, time.Now(), "two", -100, ""}

	got := balanceSet(set, 100)

	if got[0].Id != 0 {
		t.Error(fmt.Sprintf("[0] expected id 1 but got %d", got[0].Id))
	}

	if got[0].Cents != -100 {
		t.Error(fmt.Sprintf("[0] expected cents of -100, got %d", got[0].Cents))
	}

	if got[0].Balance != 200 {
		t.Error(fmt.Sprintf("[0] expected balance of 200, got %d", got[0].Balance))
	}

	if got[1].Id != 1 {
		t.Error(fmt.Sprintf("[1] expected id 1 but got %d", got[1].Id))
	}

	if got[1].Cents != -100 {
		t.Error(fmt.Sprintf("[1] expected cents of -100, got %d", got[1].Cents))
	}

	if got[1].Balance != 100 {
		t.Error(fmt.Sprintf("[1] expected balance of 100, got %d", got[1].Balance))
	}

	if len(got) != 2 {
		t.Error(fmt.Sprintf("expected 1 element but got %d", len(got)))
	}
}

func TestBalanceSetC(t *testing.T) {
	set := make([]domain.Transaction, 6)
	set[0] = domain.Transaction{1, time.Now(), "one", -10, ""}
	set[1] = domain.Transaction{1, time.Now(), "two", -20, ""}
	set[2] = domain.Transaction{1, time.Now(), "three", -30, ""}
	set[3] = domain.Transaction{1, time.Now(), "four", 40, ""}
	set[4] = domain.Transaction{1, time.Now(), "five", -50, ""}
	set[5] = domain.Transaction{1, time.Now(), "six", -60, ""}

	got := balanceSet(set, 300)

	if got[0].Cents != -10 {
		t.Error(fmt.Sprintf("[0] expected cents of -20, got %d", got[0].Cents))
	}

	if got[0].Balance != 420 {
		t.Error(fmt.Sprintf("[0] expected balance of 420, got %d", got[0].Balance))
	}

	if got[1].Cents != -20 {
		t.Error(fmt.Sprintf("[1] expected cents of -20, got %d", got[1].Cents))
	}

	if got[1].Balance != 400 {
		t.Error(fmt.Sprintf("[1] expected balance of 400, got %d", got[1].Balance))
	}

	if got[2].Cents != -30 {
		t.Error(fmt.Sprintf("[2] expected cents of -10, got %d", got[2].Cents))
	}

	if got[2].Balance != 370 {
		t.Error(fmt.Sprintf("[2] expected balance of 370, got %d", got[2].Balance))
	}

	if got[3].Cents != 40 {
		t.Error(fmt.Sprintf("[3] expected cents of -10, got %d", got[3].Cents))
	}

	if got[3].Balance != 410 {
		t.Error(fmt.Sprintf("[3] expected balance of 410, got %d", got[3].Balance))
	}

	if got[4].Cents != -50 {
		t.Error(fmt.Sprintf("[4] expected cents of -10, got %d", got[4].Cents))
	}

	if got[4].Balance != 360 {
		t.Error(fmt.Sprintf("[4] expected balance of 360, got %d", got[4].Balance))
	}

	if got[5].Cents != -60 {
		t.Error(fmt.Sprintf("[5] expected cents of -10, got %d", got[5].Cents))
	}

	if got[5].Balance != 300 {
		t.Error(fmt.Sprintf("[5] expected balance of 100, got %d", got[5].Balance))
	}

}
