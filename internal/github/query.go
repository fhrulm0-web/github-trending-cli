package github

import (
	"fmt"
	"strings"
	"time"
)

const (
	DurationDay   = "day"
	DurationWeek  = "week"
	DurationMonth = "month"
	DurationYear  = "year"

	DefaultDuration = DurationWeek
	DefaultLimit    = 10
	MinLimit        = 1
	MaxLimit        = 100
)

// ValidDurations contains all supported duration values.
var ValidDurations = map[string]bool{
	DurationDay:   true,
	DurationWeek:  true,
	DurationMonth: true,
	DurationYear:  true,
}

// CalculateSinceDate calculates the starting date based on duration and reference time in UTC.
func CalculateSinceDate(duration string, now time.Time) (time.Time, error) {
	ref := now.UTC()
	switch duration {
	case DurationDay:
		return ref.AddDate(0, 0, -1), nil
	case DurationWeek:
		return ref.AddDate(0, 0, -7), nil
	case DurationMonth:
		return ref.AddDate(0, 0, -30), nil
	case DurationYear:
		return ref.AddDate(-1, 0, 0), nil
	default:
		return time.Time{}, fmt.Errorf("invalid duration '%s'. Allowed values: day, week, month, year", duration)
	}
}

// BuildSearchQuery constructs the GitHub search query string.
func BuildSearchQuery(duration string, language string, now time.Time) (string, error) {
	since, err := CalculateSinceDate(duration, now)
	if err != nil {
		return "", err
	}

	dateStr := since.Format("2006-01-02")
	var parts []string

	cleanLang := strings.TrimSpace(language)
	if cleanLang != "" {
		parts = append(parts, fmt.Sprintf("language:%s", cleanLang))
	}

	parts = append(parts, fmt.Sprintf("created:>%s", dateStr))

	return strings.Join(parts, " "), nil
}
