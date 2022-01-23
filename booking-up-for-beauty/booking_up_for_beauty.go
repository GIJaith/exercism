package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	//fmt.Println(t)

	return t

	//panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// Current timestamp
	start := time.Now()

	// Convert date: July 25, 2019 13:45:00
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	// Find elapsed time
	elapsed := start.Sub(t)
	fmt.Println(elapsed)

	switch {
	case elapsed >= 0:
		// In the past
		return true
	case elapsed < 0:
		// In the future
		return false
	}
	return false
	//panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	switch {
	case t.Hour() >= 12 && t.Hour() < 18:
		return true
	default:
		return false
	}

	//panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	t := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
	//panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	y := time.Now().Year()
	a := fmt.Sprintf("%d-09-15", y)

	// Convert date: Mon Jan 2 15:04:05 -0700 MST 2006
	layout := "2006-01-02"
	t, _ := time.Parse(layout, a)
	return t
	//panic("Please implement the AnniversaryDate function")
}
