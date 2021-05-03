package helper

import "time"

func ParseDateString(date string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	return time.ParseInLocation(time.RFC3339, date, loc)
}
