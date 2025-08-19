package urania

import "time"

func combine(date time.Time, timestamp time.Time) time.Time {
	return time.Date(
		date.Year(), date.Month(), date.Day(),
		timestamp.Hour(), timestamp.Minute(), timestamp.Second(), timestamp.Nanosecond(),
		date.Location(),
	)
}
