package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// LogError is a utility function for logging errors in a consistent format.
func LogError(err error, context string) {
	if err != nil {
		log.Printf("ERROR: %s - %v", context, err)
	}
}

// ParseDate is a utility function for parsing date strings into time.Time.
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04", dateStr)
}

// RespondWithError is a utility function for responding with an error message.
func RespondWithError(w http.ResponseWriter, message string, code int) {
	http.Error(w, fmt.Sprintf("Error: %s", message), code)
}

// FormatRemainingTime formats the duration to a human-readable string
func FormatRemainingTime(duration time.Duration) string {
	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	return fmt.Sprintf("%d days %d hours %d minutes", days, hours, minutes)
}
