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
		err = t.Date.SetText(d[0])
		if err != nil {
			fmt.Printf("Skipping line %d: Invalid date\n", i)
			skipped += 1
			continue
		}

		err = t.Category.SetText(d[1])
		if err != nil {
			fmt.Printf("Skipping line %d: Invalid category\n", i)
			skipped += 1
			continue
		}

		err = t.Description.SetText(d[2])
		if err != nil {
			err = fmt.Errorf("Skipping line %d: %w", i, err)
			fmt.Println(err.Error())
			skipped += 1
			continue
		}

		err = t.Cents.SetText(d[3])
		if err != nil {
			fmt.Printf("Skipping line %d: Invalid amount\n", i)
			skipped += 1
			continue
		}

		if len(d) < 4 {
			err = t.Note.SetText(d[4])
			if err != nil {
				fmt.Printf("Skipping line %d: Invalid notes\n", i)
				skipped += 1
				continue
			}
		} else {
			d = append(d, "")
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
