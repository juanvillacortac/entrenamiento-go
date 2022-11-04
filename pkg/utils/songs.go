package utils

import "fmt"

func FormatDurationFromMilliseconds(milliseconds int64) string{
	var minutes int64 = (milliseconds / 1000) / 60
	var seconds int64 = (milliseconds / 1000) % 60
	return fmt.Sprintf("%d:%02d", minutes,seconds)
}
