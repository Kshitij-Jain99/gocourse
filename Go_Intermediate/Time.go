package intermediate

import (
	"fmt"
	"time"
)

func time_main() {

	// Current local time
	fmt.Println(time.Now())

	// Specific time/Date:   time.Date(year, month, day, hour, min, sec, nsec, loc)
	specificTime := time.Date(2024, time.July, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println("Specific time:", specificTime)

	// Parse time: time.Parse(layout, value) -> string to time
	//Reference: Mon Jan 2 15:04:05 MST 2006
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01")      //2020-05-01 00:00:00 +0000 UTC
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01")         //2020-05-01 00:00:00 +0000 UTC
	parsedTime2, _ := time.Parse("06-1-2", "20-5-1")             // 2020-05-01 00:00:00 +0000 UTC
	parsedTime3, _ := time.Parse("06-1-2 15-04", "20-5-1 18-03") // 2020-05-01 18:03:00 +0000 UTC
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)

	// Formatting time: Time to string
	t := time.Now()
	fmt.Println("Formatted time:", t.Format("Monday 06-01-02 15-04-05"))

	//Time manipulation
	oneDayLater := t.Add(time.Hour * 24) //Adding durations //Adds one day (24 hours).
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday()) //Extract weekday

	// Rounding and Truncating time
	fmt.Println("Rounded Time:", t.Round(time.Hour)) //Round → Rounds to nearest multiple of duration.

	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Date(2024, time.July, 8, 14, 16, 40, 00, time.UTC)

	// Convert this to the specific time zone
	tLocal := t.In(loc)

	// Perform rounding
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original Time (UTC):", t)
	fmt.Println("Original Time (Local):", tLocal)
	fmt.Println("Rounded Time (UTC):", roundedTime)
	fmt.Println("Rounded Time (Local):", roundedTimeLocal)

	fmt.Println("Truncated Time:", t.Truncate(time.Hour))

	loc2, _ := time.LoadLocation("America/New_York")

	// Convert time to location
	tInNY := time.Now().In(loc2)
	fmt.Println("New York Time:", tInNY)

	t1 := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, time.July, 4, 18, 0, 0, 0, time.UTC)
	duration := t2.Sub(t1)
	fmt.Println("Duration:", duration)

	// Compare times
	fmt.Println("t2 is after t1?", t2.After(t1)) //t1.Before(t2)
}
