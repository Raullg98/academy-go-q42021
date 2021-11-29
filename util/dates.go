package util

import (
	"regexp"
	"time"
)

// GetDateFormmated - Returns the time in  a YYYY-MM-DD format
func GetDateFormmated(date time.Time) string {
	return date.Local().Format("2006-01-02")
}

// GetDateFormmated - Validates date in a YYYY-MM-DD format
func ValidateDateFormat(date string) bool {
	regex := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	return regex.MatchString(date)
}
