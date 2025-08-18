package pannonia

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func parseDate(dayWrapper *colly.HTMLElement) time.Time {
	dateStr := dayWrapper.DOM.Find(".date").Text()
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
	for _, substring := range titlePostfixes {
		trimmed = strings.ReplaceAll(trimmed, substring, "")
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

func parseTitle(movieData *colly.HTMLElement) string {
	return trim(movieData.DOM.Find("div.title").Text())
}

func parseMovieLink(movieWrapper *colly.HTMLElement) string {
	return movieWrapper.ChildAttr("td.info a", "href")
}

func parseOriginalTitle(movieData *colly.HTMLElement) string {
	originalTitle := movieData.DOM.Find("p.title-original").Text()
	originalTitle = originalTitle[:len(originalTitle)-4]
	if originalTitle == "" {
		originalTitle = parseTitle(movieData)
	}
	return originalTitle
}

func parseYear(movieData *colly.HTMLElement) int {
	originalTitle := movieData.DOM.Find("p.title-original").Text()
	year, _ := strconv.Atoi(originalTitle[len(originalTitle)-4:])
	return year
}

func parseIsPremiere(movieWrapper *colly.HTMLElement) bool {
	return movieWrapper.DOM.Find(".premiere").Text() == "Premier"
}

func parseBookingLink(movieTime *colly.HTMLElement) string {
	return movieTime.ChildAttr("a", "href")
}

func parseIsSubbed(movieTime *colly.HTMLElement) bool {
	return movieTime.DOM.Find(".type").Text() == "F"
}

func parseDateTime(date time.Time, movieTime *colly.HTMLElement) time.Time {
	timestamp, _ := time.Parse("15:04", movieTime.DOM.Find(".time").Text())
	return combine(date, timestamp)
}
