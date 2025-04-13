package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// "7/25/2019 13:45:00"
	parsedTime, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(fmt.Sprintf("something went wrong parsing the date: %s", err.Error()))
	}
	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	parsedTime, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(fmt.Sprintf("something went wrong parsing the date: %s", err.Error()))
	}
	return now.After(parsedTime)

}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedTime, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(fmt.Sprintf("something went wrong parsing the date: %s", err.Error()))
	}
	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedTime, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(fmt.Sprintf("something went wrong parsing the date: %s", err.Error()))
	}

	formatted := parsedTime.Format("Monday, January 2, 2006, at 15:04.")
	return fmt.Sprintf("You have an appointment on %s", formatted)

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
