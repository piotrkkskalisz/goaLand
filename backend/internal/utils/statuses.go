package utils

import "time"

const (
	EditionUpcoming = "UPCOMING"
	EditionActive   = "ACTIVE"
	EditionFinished = "FINISHED"
)

func EditionStatus(startDate time.Time, endDate time.Time) string {
	now := time.Now()
	if now.Before(startDate) {
		return EditionUpcoming
	} else if now.Before(endDate) {
		return EditionActive
	} else {
		return EditionFinished
	}
}

func matchStatus(startDate time.Time, endDate time.Time) string {
	return ""
}
