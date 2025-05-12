package helper

import (
	"time"
)

// BusinessDayFromStart returns a Time object representing the day
// given x days after start, taking into account "business" days
// (that is, skipping Saturdays and Sundays)
func BusinessDayFromStart(start time.Time, days int) time.Time {
	counter := days
	currentTime := start
	for {
		currentTime = currentTime.Add(24 * time.Hour)

		switch currentTime.Weekday() {
		case time.Saturday, time.Sunday:
			continue
		}

		counter -= 1

		if counter == 0 {
			break
		}
	}

	return currentTime
}
