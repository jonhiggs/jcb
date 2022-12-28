package ui

import (
	"fmt"
	"jcb/config"
	"strings"

	"github.com/gdamore/tcell/v2"
)

var findQuery string

func handleFind(ev *tcell.EventKey) *tcell.EventKey {
	switch ev.Rune() {
	case '/':
		openPrompt("/", "", func(ev *tcell.EventKey) *tcell.EventKey {
			panels.HidePanel("prompt")
			findQuery = promptInputField.GetText()
			selectNextMatch()
			return nil
		})
	case '?':
		openPrompt("?", "", func(ev *tcell.EventKey) *tcell.EventKey {
			panels.HidePanel("prompt")
			findQuery = promptInputField.GetText()
			selectPrevMatch()
			return nil
		})
	case 'T':
		openPrompt("Tag matched transactions:", selectedDescription(), func(ev *tcell.EventKey) *tcell.EventKey {
			panels.HidePanel("prompt")
			findQuery = promptInputField.GetText()
			tagMatches()
			return nil
		})
	}

	return nil
}

func handleFindForwards(ev *tcell.EventKey) *tcell.EventKey {
	selectNextMatch()
	return nil
}

func handleFindBackwards(ev *tcell.EventKey) *tcell.EventKey {
	selectPrevMatch()
	return nil
}

func selectNextMatch() {
	curRow, _ := transactionsTable.GetSelection()

	for i := curRow + 1; i != curRow; i++ {
		if strings.Contains(strings.ToLower(transactionsTable.GetCell(i, config.DESCRIPTION_COLUMN).GetText()), strings.ToLower(findQuery)) {
			transactionsTable.Select(i, 0)
			return
		}

		if i == len(transactionIds) {
			i = 0
		}
	}

	printStatus("No matches found")
}

func selectPrevMatch() {
	curRow, _ := transactionsTable.GetSelection()

	for i := curRow - 1; i != curRow; i-- {
		if strings.Contains(strings.ToLower(transactionsTable.GetCell(i, config.DESCRIPTION_COLUMN).GetText()), strings.ToLower(findQuery)) {
			transactionsTable.Select(i, 0)
			break
		}

		if i == 0 {
			i = len(transactionIds) - 1
		}
	}

	printStatus("No matches found")
}

func tagMatches() {
	curRow, _ := transactionsTable.GetSelection()

	matchCount := 0

	for i := 1; i < len(transactionIds); i++ {
		if isCommitted(i) || isTagged(i) {
			continue
		}

		if strings.Contains(strings.ToLower(transactionsTable.GetCell(i, config.DESCRIPTION_COLUMN).GetText()), strings.ToLower(findQuery)) {
			matchCount += 1
			applyTag(i)
		}
	}

	transactionsTable.Select(curRow, 0)

	printStatus(fmt.Sprintf("Tagged %d transactions", matchCount))
	updateTransactionsTable()
}
