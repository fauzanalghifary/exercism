// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	nextOneGigaSecond := t.Unix() + 1_000_000_000
	nextDate := time.Unix(nextOneGigaSecond, 0)
	return nextDate
}

// Community Solution
//func AddGigasecond(t time.Time) time.Time {
//	return t.Add(time.Second * 1000000000)
//}
