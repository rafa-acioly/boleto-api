package util

import (
	"time"
)

func Duration(callback func()) (duration time.Duration) {
	start := time.Now()
	callback()
	end := time.Now()
	duration = end.Sub(start)
	return
}
