package transaction

func UpdateCategory(s string, ts []*Transaction) []*Transaction {
	var modifedTransactions []*Transaction

	for _, t := range ts {
		if t.Category.GetText() != s {
			t.Category.SetText(s)
			modifedTransactions = append(modifedTransactions, t)
		}
	}

	return modifedTransactions
}

func UpdateDescription(s string, ts []*Transaction) []*Transaction {
	var modifedTransactions []*Transaction

	for _, t := range ts {
		if t.Description.GetText() != s {
			t.Description.SetText(s)
			modifedTransactions = append(modifedTransactions, t)
		}
	}

	return modifedTransactions
}

func UpdateDate(s string, ts []*Transaction) []*Transaction {
	var modifedTransactions []*Transaction

	// TODO: prevent data from being set too early

	for _, t := range ts {
		if t.Date.GetText() != s {
			t.Date.SetText(s)
			modifedTransactions = append(modifedTransactions, t)
		}
	}

	return modifedTransactions
}

func UpdateCents(s string, ts []*Transaction) []*Transaction {
	var modifedTransactions []*Transaction

	// TODO: prevent data from being set too early

	for _, t := range ts {
		if t.Cents.GetText() != s {
			t.Cents.SetText(s)
			modifedTransactions = append(modifedTransactions, t)
		}
	}

	return modifedTransactions
}
