/**
Go supports time formatting and parsing via pattern-based layouts.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	print := fmt.Println

	// Here’s a basic example of formatting a time according to RFC3339, using the corresponding layout constant.
	now := time.Now()
	print(now.Format(time.RFC3339))

	// Time parsing uses the same layout values as Format.
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	print(t1)

	// Format and Parse use example-based layouts. Usually you’ll use a constant from time for these layouts, but you can also supply custom layouts.
	// Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006 to show the pattern with which to format/parse a given time/string.
	// The example time must be exactly as shown: the year 2006, 15 for the hour, Monday for the day of the week, etc.
	print(now.Format("3:04PM"))
	print(now.Format("Mon Jan _2 15:04:05 2006"))
	print(now.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	print(t2)

	// For purely numeric representations you can also use standard string formatting with the extracted components of the time value.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	// Parse will return an error on malformed input explaining the parsing problem.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	print(e)
}

// output

/**
2020-01-31T20:38:21+08:00
2012-11-01 22:08:41 +0000 +0000
8:38PM
Fri Jan 31 20:38:21 2020
2020-01-31T20:38:21.907191+08:00
0000-01-01 20:41:00 +0000 UTC
2020-01-31T20:38:21-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
*/
