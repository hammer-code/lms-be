package utils

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

func ParseDate(dateStr string) (null.Time, error) {
	if dateStr == "" {
		return null.Time{}, nil // If no date provided, return a null time
	}

	// Define the expected date format (e.g., "2006-01-02" for YYYY-MM-DD)
	layout := "2006-01-02"

	// Parse the date string
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return null.Time{}, err // Return an error if the format is wrong
	}

	// Return the parsed time as a non-null time
	return null.NewTime(parsedTime, true), nil
}
