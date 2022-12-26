package category

import (
	"jcb/db"
	"time"
)

func All() []string {
	var categories []string
	for _, name := range db.Categories() {
		if name == "" {
			name = "uncategorised"
		}

		categories = append(categories, name)
	}

	return categories
}

func Sum(category string, startTime time.Time, endTime time.Time) int64 {
	if category == "uncategorised" {
		category = ""
	}
	return db.CategorySum(category, startTime, endTime)
}
