package tsv

import (
	"bufio"
	"fmt"
	"jcb/lib/transaction"
	"log"
	"os"
	"strings"
)

func Import(f string) bool {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	imported := 0
	skipped := 0
	for scanner.Scan() {
		i += 1
		d := strings.Split(scanner.Text(), "\t")

		if len(d) < 3 {
			fmt.Printf("Skipping line %d: Expected at least 4 columns but got %d\n", i, len(d))
			skipped += 1
			continue
		}

		t := new(transaction.Transaction)
		err := t.SetText(d)
		if err != nil {
			fmt.Printf("Skipping line %d: %T\n", i)
			skipped += 1
			continue
		}

		if !t.IsUniq() {
			fmt.Printf("Skipping line %d: Transaction is not unique\n", i)
			skipped += 1
			continue
		}

		t.Save()
		imported += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nImported %d lines. Skipped %d lines.\n", imported, skipped)

	return true
}
