package gobilla

import (
	"time"
)

// Date format constants.
const (
	RFC1123GMT    = "Mon, 02 Jan 2006 15:04:05 GMT"
	ShortDate     = "20060102"
	ShortDateTime = "20060102T150405Z"
)

func getRFC1123GMT(now time.Time) string {
	return now.UTC().Format(RFC1123GMT)
}

func getShortDate(now time.Time) string {
	return now.UTC().Format(ShortDate)
}

func getShortDateTime(now time.Time) string {
	return now.UTC().Format(ShortDateTime)
}
