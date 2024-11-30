package utils

import "time"

const WEEK = time.Hour * 24 * 14

func TwoWeeksFromNow() time.Time {
	return time.Now().Add(2 * WEEK)
}
