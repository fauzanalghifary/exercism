package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	day := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	if wSched == Teenth {
		for i := 0; i < 7; i++ {
			finalDay := day.AddDate(0, 0, 12+i)
			if finalDay.Weekday() == wDay {
				return finalDay.Day()
			}
		}
	}

	var foundDay []int
	for i := 0; i < 31; i++ {
		finalDay := day.AddDate(0, 0, i)
		if finalDay.Weekday() == wDay && finalDay.Month() == day.Month() {
			foundDay = append(foundDay, finalDay.Day())
		}
	}

	if wSched < 4 {
		return foundDay[wSched]
	}

	return foundDay[len(foundDay)-1]
}
