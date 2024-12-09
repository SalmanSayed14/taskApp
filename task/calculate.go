package task

import (
	"fmt"
	"time"
)

func CalculateRemainingTime(deadline time.Time) string {
	now := time.Now()
	if deadline.Before(now) {
		return "Expired"
	}

	duration := deadline.Sub(now)
	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	return fmt.Sprintf("%d days %d hours %d minutes", days, hours, minutes)
}

func CalculateRemainingDuration(deadline time.Time) time.Duration {
	now := time.Now()
	if deadline.Before(now) {
		return -1
	}
	return deadline.Sub(now)
}
