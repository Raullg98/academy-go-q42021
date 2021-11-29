package util

import (
	"regexp"
	"time"
)

func GetDateFormmated(date time.Time) string {
	return date.Local().Format("2006-01-02")
}

func ValidateDateFormat(date string) bool {
	regex := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	return regex.MatchString(date)
}
