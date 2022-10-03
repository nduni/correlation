package utils

import (
	"time"
)

const DATE_TIME_LAYOUT = "2006-01-02T15:04"

func ParseDateTimeToUTCString(dateTime, timezone string) (string, error) {
	loc, _ := time.LoadLocation(timezone)
	dT, err := time.ParseInLocation(DATE_TIME_LAYOUT, dateTime, loc)
	if err != nil {
		return "", err
	}
	return dT.UTC().Format(DATE_TIME_LAYOUT), nil
}
