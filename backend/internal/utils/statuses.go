package utils

import "time"

const (
	EditionUpcoming = "UPCOMING"
	EditionActive   = "ACTIVE"
	EditionFinished = "FINISHED"
)

var UpcomingMatchStatuses = []string{
	"SCHEDULED",
	"TIMED",
}

var LiveMatchStatuses = []string{
	"LIVE",
	"IN_PLAY",
	"PAUSED",
	"EXTRA_TIME",
	"PENALTY_SHOOTOUT",
}

var FinishedMatchStatuses = []string{
	"FINISHED",
	"AWARDED",
}

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

func IsCurrent(status string) bool {
	return status == EditionActive || status == EditionUpcoming
}
