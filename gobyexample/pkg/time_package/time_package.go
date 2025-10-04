package timepackage

// Go offers extensive support for times and durations; here are some examples.

import (
	"fmt"
	"time"
)

func Run() {

	fmt.Println("\nTime Package: ")

	p := fmt.Println

	// We’ll start by getting the current time.

	now := time.Now()
	p(now)

	// You can build a time struct by providing the year, month, day, etc. Times are always associated with a Location, i.e. time zone.

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// You can extract the various components of the time value as expected.

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// The Monday-Sunday Weekday is also available.

	p(then.Weekday())

	// These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// The Sub methods returns a Duration representing the interval between two times.

	diff := now.Sub(then)
	p(diff)

	// We can compute the length of the duration in various units.

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.

	p(then.Add(diff))
	p(then.Add(-diff))
}

// Time Package:
// 2025-10-04 15:02:58.847032 +0200 CEST m=+0.000197876
// 2009-11-17 20:34:58.651387237 +0000 UTC
// 2009
// November
// 17
// 20
// 34
// 58
// 651387237
// UTC
// Tuesday
// true
// false
// false
// 139192h28m0.195644763s
// 139192.46672101243
// 8.351548003260746e+06
// 5.0109288019564474e+08
// 501092880195644763
// 2025-10-04 13:02:58.847032 +0000 UTC
// 1994-01-01 04:06:58.455742474 +0000 UTC
