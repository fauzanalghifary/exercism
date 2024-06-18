package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"

	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	now := time.Now()
	appointment, _ := time.Parse(layout, date)
	return now.After(appointment)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return 12 <= t.Hour() && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	//return fmt.Sprintf(
	//	"You have an appointment on %s, %s %d, %d, at %d:%d.",
	//	t.Weekday(),
	//	t.Month(),
	//	t.Day(),
	//	t.Year(),
	//	t.Hour(),
	//	t.Minute(),
	//)
	return fmt.Sprintf(
		"You have an appointment on %s",
		t.Format("Monday, January 2, 2006, at 15:04."),
	)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	//return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
	t, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return t
}
