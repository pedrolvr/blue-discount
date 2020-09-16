package util

import "time"

func IsBirthday(date time.Time) bool {
	_, month, day := time.Now().Date()
	return date.Day() == day && date.Month() == month
}

func DateIsToday(date time.Time) bool {
	year, month, day := time.Now().Date()
	return date.Day() == day && date.Month() == month && date.Year() == year
}
