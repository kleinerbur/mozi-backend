package pannonia

import (
	"fmt"
	"strings"
	"time"
)

func parseDate(dateStr string) time.Time {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if dateStr == "Ma" {
		return today
	}
	if dateStr == "Holnap" {
		return today.AddDate(0, 0, 1)
	}
	monthTranslations := map[string]string{
		"január":     "jan",
		"február":    "feb",
		"március":    "mar",
		"április":    "apr",
		"május":      "may",
		"június":     "jun",
		"július":     "jul",
		"augusztus":  "aug",
		"szeptember": "sep",
		"október":    "oct",
		"november":   "nov",
		"december":   "dec",
	}
	for hun, eng := range monthTranslations {
		dateStr = strings.ReplaceAll(dateStr, hun, eng)
	}
	date, err := time.ParseInLocation("Jan 2.", dateStr, today.Location())
	if err != nil {
		fmt.Println(err)
	}
	date = date.AddDate(today.Year(), 0, 0)
	if date.Before(today) {
		date = date.AddDate(1, 0, 0)
	}
	return date
}

func combine(date time.Time, timestamp time.Time) time.Time {
	return time.Date(
		date.Year(), date.Month(), date.Day(),
		timestamp.Hour(), timestamp.Minute(), timestamp.Second(), timestamp.Nanosecond(),
		date.Location(),
	)
}

func trim(title string) string {
	trimmed := title
	for _, delimiter := range titlePostfixes {
		trimmed = strings.Split(trimmed, delimiter)[0]
	}
	return trimmed
}

func isAnActualMovie(title string) bool {
	for _, keyword := range excludedKeywords {
		if strings.Contains(title, keyword) {
			return false
		}
	}
	return true
}
