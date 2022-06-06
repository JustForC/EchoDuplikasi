package helpers

import "time"

func ConvertTime(t string) time.Time {
	date, _ := time.Parse("2006-01-02", t)

	return date
}
