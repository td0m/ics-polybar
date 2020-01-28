package main

import (
	"fmt"
	"github.com/apognu/gocal"
	"os"
	"strings"
	"time"
)

func prettyTime(now time.Time, date *time.Time) string {
	dayDiff := date.Day() - now.Day()
	if dayDiff > 0 {
		return fmt.Sprintf("at %d:%02d", date.Hour(), date.Minute())
	}
	hourDiff := date.Hour() - now.Hour()
	if hourDiff <= 1 {
		minuteDiff := 60*hourDiff + date.Minute() - now.Minute()
		return fmt.Sprintf("in %d minutes", minuteDiff)

	}
	return fmt.Sprintf("in %d hours", hourDiff)
}

func prettyDay(now time.Time, date *time.Time) string {
	dayDiff := date.Day() - now.Day()
	if dayDiff == 0 {
		return ""
	}
	if dayDiff == 1 {
		return "tomorrow"
	}
	return fmt.Sprintf("in %d days", now.Day())

}

func prettyDate(now time.Time, date *time.Time) string {
	return fmt.Sprintf("%s %s", prettyDay(now, date), prettyTime(now, date))
}

func prettyName(summary string) string {
	parts := strings.Split(summary, "-")
	if len(parts) == 1 {
		return summary
	}
	return strings.TrimSpace(parts[1])
}

func main() {
	f, _ := os.Open("./calendar.ics")
	defer f.Close()

	start, end := time.Now(), time.Now().Add(12*30*24*time.Hour)

	c := gocal.NewParser(f)
	c.Start, c.End = &start, &end
	c.Parse()

	now := time.Now()
	// now = now.Add(time.Hour * 24)

	var next gocal.Event
	for _, e := range c.Events {
		// fmt.Printf("%s on %s by\n", e.Summary, e.Start)
		if e.Start.After(now) {
			next = e
			break
		}
	}
	fmt.Printf("%s, %s\n", prettyName(next.Summary), prettyDate(now, next.Start))
}
