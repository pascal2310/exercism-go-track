// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

const thousand = 1000
const million = 1000000
const gigasecond = thousand * million * time.Second

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
