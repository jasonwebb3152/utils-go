package dates

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05" // format for mysql datetime
)

/* This is how we will store the time as a string in Go and
we can formate the string by passing in a placeholder as seen below. */

func GetNow() time.Time {
	return time.Now().UTC() // Standard time is the best to use (because cloud service) so we call .UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
