package date

import (
	"time"
)

// Date format constants.
const (
	RFC1123GMT    = "Mon, 02 Jan 2006 15:04:05 GMT"
	ShortDate     = "20060102"
	ShortDateTime = "20060102T150405Z"
)

// GetRFC1123GMT returns now in RFC1123GMT format.
func GetRFC1123GMT(now time.Time) string {
	return now.UTC().Format(RFC1123GMT)
}

// GetShortDate returns now in short date format.
func GetShortDate(now time.Time) string {
	return now.UTC().Format(ShortDate)
}

// GetShortDateTime returns now in short date time format.
func GetShortDateTime(now time.Time) string {
	return now.UTC().Format(ShortDateTime)
}
