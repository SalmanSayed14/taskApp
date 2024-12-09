package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func LogError(err error, context string) {
	if err != nil {
		log.Printf("ERROR: %s - %v", context, err)
	}
}

func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04", dateStr)
}

func RespondWithError(w http.ResponseWriter, message string, code int) {
	http.Error(w, fmt.Sprintf("Error: %s", message), code)
}

func FormatRemainingTime(duration time.Duration) string {
	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	return fmt.Sprintf("%d days %d hours %d minutes", days, hours, minutes)
}
