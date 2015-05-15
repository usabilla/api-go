package gobilla

import (
	"time"
)

// Date format constants.
const (
	rfc1123gmt    = "Mon, 02 Jan 2006 15:04:05 GMT"
	shortDate     = "20060102"
	shortDateTime = "20060102T150405Z"
)

// getRFC1123GMT returns now in RFC1123GMT format.
func getRFC1123GMT(now time.Time) string {
	return now.UTC().Format(rfc1123gmt)
}

// getShortDate returns now in short date format.
func getShortDate(now time.Time) string {
	return now.UTC().Format(shortDate)
}

// getShortDateTime returns now in short date time format.
func getShortDateTime(now time.Time) string {
	return now.UTC().Format(shortDateTime)
}
