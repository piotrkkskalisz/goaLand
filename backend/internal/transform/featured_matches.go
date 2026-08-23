package transform

import (
	"backend/internal/database"
	"backend/internal/utils"
	"slices"
	"time"
)

const daysBefore = 3
const daysAfter = 6

func dayStart(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func From() time.Time {
	return dayStart(time.Now()).AddDate(0, 0, -daysBefore)
}

func To() time.Time {
	return dayStart(time.Now()).AddDate(0, 0, daysAfter+1)
}

func chooseFeatured(matches []database.Match, now time.Time) []database.Match {
	today := dayStart(now)
	tomorrow := today.AddDate(0, 0, 1)
	featured := make([]database.Match, 0)

	earlierMatchIndex := -1
	afterMatchIndex := -1

	hasFinishedToday := false
	hasUpcomingToday := false

	for i, match := range matches {
		if match.StartTime.Before(today) && slices.Contains(utils.FinishedMatchStatuses, match.Status) {
			earlierMatchIndex = i
			continue
		}
		if !match.StartTime.Before(tomorrow) {
			afterMatchIndex = i
			break
		}

		featured = append(featured, match)
		hasFinishedToday = hasFinishedToday || slices.Contains(utils.FinishedMatchStatuses, match.Status)
		hasUpcomingToday = hasUpcomingToday || slices.Contains(utils.UpcomingMatchStatuses, match.Status)
	}
	if !hasFinishedToday && earlierMatchIndex > -1 {
		featured = append([]database.Match{matches[earlierMatchIndex]}, featured...)
	}
	if !hasUpcomingToday && afterMatchIndex > -1 {
		featured = append(featured, matches[afterMatchIndex])
		if afterMatchIndex+1 < len(matches) && len(featured) < 3 {
			featured = append(featured, matches[afterMatchIndex+1])
		}
	}

	return featured
}
func GetFeaturedMatches(matches []database.Match) []MatchResponse {
	matches = chooseFeatured(matches, time.Now())

	matchResponse := make([]MatchResponse, 0, len(matches))
	for _, match := range matches {
		matchResponse = append(matchResponse, createMatchResponse(match))
	}
	return matchResponse
}
