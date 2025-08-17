package artmozi

import "time"

func ScheduleWeek(date time.Time) int {
	_, week := date.AddDate(0, 0, -3).ISOWeek()
	return week
}
