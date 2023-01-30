package helper

import (
	"fmt"
)

var hr, min, sec int

func GetAppRunTime() string {
	sec = sec + 1
	if sec == 60 {
		sec = 0
		min = min + 1
	}
	if min == 60 {
		min = 0
		hr = hr + 1
	}
	return fmt.Sprintf("%02d:%02d:%02d", hr, min, sec)
}
