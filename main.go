package main

import (
	"fmt"
	"github.com/apognu/gocal"
	"math"
	"os"
	"strings"
	"time"
)

// returns a formatted date string
func formatDate(date time.Time) string {
	now := time.Now()
	diff := date.Sub(now)
	minuteDiff := int(math.Floor(diff.Minutes())) % 60
	hourDiff := int(math.Floor(diff.Hours()))

	if diff.Minutes() <= 60 {
		return fmt.Sprintf("in %d minutes", minuteDiff)
	}
	if diff.Hours() <= 12 {
		return fmt.Sprintf("in %dh %d min", hourDiff, minuteDiff)
	}
	formattedHour := date.Format("15:04")
	if now.YearDay() == date.YearDay() {
		return fmt.Sprintf("today at %s", formattedHour)
	}
	if now.YearDay()+1 == date.YearDay() {
		return fmt.Sprintf("tomorrow at %s", formattedHour)
	}
	if now.YearDay()+7 >= date.YearDay() {
		return date.Format("Mon at 15:04")
	}

	return date.Format("Mon 2 Jan at 15:04")
}

// returns a formatted name (you might have to change this)
// my events follow a "{MODULE_CODE} - {MODULE_NAME}" convention
// returns whole string if convention not recognised
func formatName(summary string) string {
	parts := strings.Split(summary, "-")
	if len(parts) != 2 {
		return summary
	}
	return strings.TrimSpace(parts[1])
}

func main() {
	path := strings.Join(os.Args[1:], " ")
	if len(path) == 0 {
		panic("Please provide the path to the config file")
	}
	// open the calendar file
	f, err := os.Open(path)
	if err != nil {
		panic("Failed to read config file, located in '" + path + "'")
	}
	defer f.Close()

	start, end := time.Now(), time.Now().Add(12*30*24*time.Hour)

	// parse the ics file
	c := gocal.NewParser(f)
	c.Start, c.End = &start, &end
	c.Parse()

	// get the next event in the calendar
	now := time.Now()
	var next gocal.Event
	for _, e := range c.Events {
		// fmt.Printf("%s on %s by\n", e.Summary, e.Start)
		if e.Start.After(now) {
			next = e
			break
		}
	}
	// output the formatted output to stdout
	// if you would like to add a room, do it here
	// most of my lectures are in the same theathre anyway, so I don't need to waste space on a room number
	fmt.Printf("%s, %s\n", formatName(next.Summary), formatDate(*next.Start))
}
