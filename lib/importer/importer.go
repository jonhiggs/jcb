package importer

import (
	"bufio"
	"fmt"
	"jcb/domain"
	dataf "jcb/lib/formatter/data"
	"jcb/lib/transaction"
	"jcb/lib/validator"
	"log"
	"os"
	"strings"
)

func Tsv(f string) bool {
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

		if validator.Date(d[0]) != nil {
			fmt.Printf("Skipping line %d: Invalid date\n", i)
			skipped += 1
			continue
		}

		if validator.Category(d[1]) != nil {
			fmt.Printf("Skipping line %d: Invalid category\n", i)
			skipped += 1
			continue
		}

		if validator.Description(d[2]) != nil {
			fmt.Printf("Skipping line %d: Invalid description\n", i)
			skipped += 1
			continue
		}

		if validator.Cents(d[3]) != nil {
			fmt.Printf("Skipping line %d: Invalid amount\n", i)
			skipped += 1
			continue
		}

		if len(d) < 5 {
			if validator.Notes(d[5]) != nil {
				fmt.Printf("Skipping line %d: Invalid notes\n", i)
				skipped += 1
				continue
			}
		} else {
			d = append(d, "")
		}

		t := domain.Transaction{
			-1,
			dataf.Date(d[0]),
			dataf.Description(d[2]),
			dataf.Cents(d[3]),
			dataf.Notes(d[4]),
			dataf.Category(d[1]),
		}

		if !transaction.Uniq(t) {
			fmt.Printf("Skipping line %d: Transaction is not unique\n", i)
			skipped += 1
			continue
		}

		_, err := transaction.Insert(t)
		if err != nil {
			fmt.Printf("Failed to import line %d: %s\n", i, err)
			continue
		}

		imported += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nImported %d lines. Skipped %d lines.\n", imported, skipped)

	return true
}
