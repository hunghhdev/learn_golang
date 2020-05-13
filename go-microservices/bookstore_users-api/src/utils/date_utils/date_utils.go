package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// GetNow return time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString return string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat return string
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
